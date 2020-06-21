import React, {useEffect, useState} from "react";
import {List, message, Skeleton, Tabs} from "antd";
import ArticleList from "../../component/articleList/articleList";
import {RouteComponentProps} from "react-router-dom";
import {CurrentUserQuery, useArticlesQuery, UserEdge, useUsersLazyQuery} from "../../generated/graphql";
import {QueryResult} from "@apollo/react-common";
import InfiniteScroll from "react-infinite-scroller";
import UserInfo from "../../component/userInfo/userInfo";
import './index.less'

export default function SearchResult(props: RouteComponentProps & { currentUser: CurrentUserQuery | undefined }) {

    const {data, error, loading, fetchMore, refetch} = useArticlesQuery({
        // @ts-ignore
        variables: {condition: props.location.state.q},
        fetchPolicy: "no-cache",
    })

    const [fetchUsers, {data: users, error: usersErr, loading: usersLoading, fetchMore: fetchMoreUsers}] = useUsersLazyQuery()

    const [list, setList] = useState()
    const [userList, setUserList] = useState()
    const [hasMore, setHasMore] = useState(false)
    const [key, setKey] = useState('article')

    useEffect(() => {
        document.title = '搜索'
    }, [])

    useEffect(() => {
        if (error) {
            message.error(error + '')
        }
        if (data && data.Articles.edges) {
            setList(data.Articles.edges)
        } else {
            setList([])
        }
    }, [data, error])

    useEffect(() => {
        if (usersErr) {
            message.error(usersErr + '')
        }
        if (users && users.Users.edges) {
            setUserList(users.Users.edges)
        } else {
            setUserList([])
        }
    }, [users, usersErr])

    useEffect(() => {
        if (key === 'article') {
            // @ts-ignore
            refetch({condition: props.location.state.q})
        } else {
            // @ts-ignore
            fetchUsers({variables: {username: props.location.state.q}})
        }
        // @ts-ignore
    }, [fetchUsers, key, props.location.state.q, refetch])

    const fetchData = () => {
        if (key === 'article') {
            fetchMore({
                // @ts-ignore
                updateQuery: ({fetchMoreResult}: { fetchMoreResult: QueryResult }) => {
                    const newEdges = fetchMoreResult.data.Articles.edges;
                    const pageInfo = fetchMoreResult.data.Articles.pageInfo;
                    setList(list.concat(...newEdges))
                    setHasMore(pageInfo.hasNextPage)
                }
            },).catch((reason: any) => message.error(reason + ''))
        } else {
            fetchMoreUsers({
                // @ts-ignore
                updateQuery: ({fetchMoreResult}: { fetchMoreResult: QueryResult }) => {
                    const newEdges = fetchMoreResult.data.Users.edges;
                    const pageInfo = fetchMoreResult.data.Users.pageInfo;
                    setUserList(userList.concat(...newEdges))
                    setHasMore(pageInfo.hasNextPage)
                }
            },).catch((reason: any) => message.error(reason + ''))
        }
    }

    const onKeyChange = (key: string) => {
        setKey(key)
    }


    return (
        <Tabs tabPosition={"left"} defaultActiveKey={'article'} onChange={onKeyChange}>
            <Tabs.TabPane tab='文章' key='article'>
                <ArticleList curId={props.currentUser?.CurrentUser.id}
                             fetchData={fetchData}
                             loading={loading}
                             hasMore={hasMore}
                             data={list}/>
            </Tabs.TabPane>
            <Tabs.TabPane tab='用户' key='user'>
                <InfiniteScroll
                    initialLoad={false}
                    pageStart={0}
                    loadMore={fetchData}
                    hasMore={!usersLoading && hasMore}>
                    <List
                        itemLayout="vertical"
                        size="small"
                        loading={usersLoading}
                        dataSource={userList}
                        renderItem={(item: UserEdge) => {
                            if (item.node.id !== props.currentUser?.CurrentUser.id) {
                                return (
                                    <List.Item
                                        key={item.node.id}
                                    >
                                        <Skeleton loading={usersLoading} active>
                                            <UserInfo loading={usersLoading}
                                                      data={{User: item.node}}
                                                      currentUserId={props.currentUser?.CurrentUser.id}
                                                      className={'search-card'}
                                            />
                                        </Skeleton>
                                    </List.Item>
                                )
                            } else {
                                return (<div/>)
                            }
                        }}
                    />
                </InfiniteScroll>
            </Tabs.TabPane>
        </Tabs>
    )
}