import React, {useEffect, useState} from "react";
import {Col, Empty, message, Row, Tabs} from "antd";
import {RouteComponentProps} from "react-router-dom";
import {
    CurrentUserQuery,
    useArticlesLazyQuery,
    useUserQuery
} from "../../generated/graphql";
import ResultPage from "../../component/result/result";
import UserInfo from "../../component/userInfo/userInfo";
import ArticleList from "../../component/articleList/articleList";
import {QueryResult} from "@apollo/react-common";
import Introduce from "../../component/userInfo/introduce";

const {TabPane} = Tabs

export default function UserCenter(props: RouteComponentProps & { currentUser: CurrentUserQuery | undefined }) {

    // @ts-ignore
    const {data, loading, error} = useUserQuery({variables: {id: parseInt(props.match.params.id)}})
    const [articles, {data: articleList, error: articlesError, loading: articlesLoading, fetchMore}] = useArticlesLazyQuery()

    const [list, setList] = useState()
    const [hasMore, setHasMore] = useState(false)

    useEffect(() => {
        if (data) {
            document.title = data.User.username
            setList([])
            articles({variables: {uid: data.User.id, cursor: null}})
        }
    }, [data, articles])

    useEffect(() => {
        if (articlesError) {
            message.error(articlesError + '')
        }
        if (articleList && articleList.Articles.edges) {
            console.log(articleList.Articles.edges)
            setList(articleList.Articles.edges)
        }
    }, [articleList, articlesError])

    const fetchData = () => {
        if (data) {
            fetchMore({
                variables: {uid: data.User.id},
                // @ts-ignore
                updateQuery: ({fetchMoreResult}: { fetchMoreResult: QueryResult }) => {
                    const newEdges = fetchMoreResult.data.HotArticles.edges;
                    const pageInfo = fetchMoreResult.data.HotArticles.pageInfo;
                    setList(list.concat(...newEdges))
                    setHasMore(pageInfo.hasNextPage)
                }
            },).catch((reason: any) => message.error(reason + ''))
        }
    }

    if (error) {
        return (<ResultPage status={"error"} title={"打开个人中心错误"} subTitle={""}/>)
    }

    return (
        <Row>
            <Col xs={24} sm={24} md={18} lg={18} xl={16} xxl={16}>
                <UserInfo loading={loading} data={data} currentUserId={props.currentUser?.CurrentUser.id}/>
                <Tabs defaultActiveKey="article" tabBarStyle={{display: "flex", paddingLeft: 50, fontSize: 16}}>
                    <TabPane key='article' tab='文章'>
                        <ArticleList
                            curId={props.currentUser?.CurrentUser.id}
                            fetchData={fetchData}
                            loading={articlesLoading && list}
                            hasMore={hasMore}
                            data={list}
                            locate={{emptyText: <Empty description="还没发表文章哦😦"/>}}
                        />
                    </TabPane>
                </Tabs>
            </Col>
            <Col xs={0} sm={0} md={6} lg={6} xl={8} xxl={8}>
                <Introduce data={data} currentUser={props.currentUser}/>
            </Col>
        </Row>
    )
}