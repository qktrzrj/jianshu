import React, {useEffect, useState} from "react";
import {RouteComponentProps} from "react-router-dom";
import {CurrentUserQuery, useArticlesLazyQuery, useFollowListQuery} from "../../generated/graphql";
import ResultPage from "../../component/result/result";
import {Avatar, Empty, message, Skeleton, Tabs} from "antd";
import ArticleList from "../../component/articleList/articleList";
import {QueryResult} from "@apollo/react-common";
import UserInfo from "../../component/userInfo/userInfo";

export default function Subscriptions(props: RouteComponentProps & { currentUser: CurrentUserQuery }) {

    const {data: followList, loading, error} = useFollowListQuery({variables: {id: props.currentUser.CurrentUser.id}})
    const [fetchArticles, {data: articles, loading: loadArticles, error: fetchArtcleErr, fetchMore}] = useArticlesLazyQuery()

    const [list, setList] = useState()
    const [hasMore, setHasMore] = useState(false)
    const [key, setKey] = useState(0)

    useEffect(() => {
        if (followList?.Followed && !list) {
            fetchArticles({variables: {uid: followList.Followed[key].id, cursor: null}})
        }
    }, [list, followList, fetchArticles, key])

    useEffect(() => {
        if (!list && articles) {
            setList(articles.Articles.edges)
        }
    }, [list, articles])

    useEffect(() => {
        if (fetchArtcleErr) {
            message.error(fetchArtcleErr.message)
        }
    }, [fetchArtcleErr])


    const fetchData = () => {
        fetchMore({
            // @ts-ignore
            updateQuery: ({fetchMoreResult}: { fetchMoreResult: QueryResult }) => {
                const newEdges = fetchMoreResult.data.HotArticles.edges;
                const pageInfo = fetchMoreResult.data.HotArticles.pageInfo;
                setList(list.concat(...newEdges))
                setHasMore(pageInfo.hasNextPage)
            }
        }).catch((reason: any) => message.error(reason + ''))
    }

    const onChange = (activeKey: string) => {
        setList(undefined)
        if (followList?.Followed) {
            fetchArticles({variables: {uid: followList.Followed[parseInt(activeKey)].id}})
        }
        setKey(parseInt(activeKey))
    }

    if (error) {
        return <ResultPage status={"error"} title="获取关注列表失败" subTitle={error.message}/>
    }

    return (
        <Skeleton loading={loading}>
            {followList?.Followed &&
            <Tabs defaultActiveKey={'0'}
                  tabPosition={"left"}
                  type={"card"}
                  onChange={onChange}
            >
                {followList.Followed.map((item, index) => {
                    return (
                        <Tabs.TabPane key={index + ''}
                                      tab={<span><Avatar style={{marginRight: 5}}
                                                         src={item.avatar}/>{item.username}</span>}>
                            <UserInfo loading={false} data={{User: item}}
                                      currentUserId={props.currentUser.CurrentUser.id}/>
                            <ArticleList
                                curId={props.currentUser.CurrentUser.id}
                                fetchData={fetchData}
                                loading={loadArticles}
                                hasMore={hasMore}
                                data={list || []}
                                locate={{emptyText: <Empty description="尚未发表文章"/>}}
                            />
                        </Tabs.TabPane>
                    )
                })}
            </Tabs>}
            {!followList?.Followed &&
            <Empty description="还没有关注任何人哦"/>
            }
        </Skeleton>
    )
}