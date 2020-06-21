import React, {useEffect, useState} from "react";
import {Col, Empty, message, Row, Tabs} from "antd";
import {RouteComponentProps} from "react-router-dom";
import {
    CurrentUserQuery,
    useLikeArticlesQuery,
} from "../../../generated/graphql";
import UserInfo from "../../../component/userInfo/userInfo";
import ArticleList from "../../../component/articleList/articleList";
import {QueryResult} from "@apollo/react-common";
import Introduce from "../../../component/userInfo/introduce";

const {TabPane} = Tabs

export default function UserLike(props: RouteComponentProps & { currentUser: CurrentUserQuery | undefined }) {

    const {data: articleList, error: articlesError, loading: articlesLoading, fetchMore} = useLikeArticlesQuery()

    const [list, setList] = useState()
    const [hasMore, setHasMore] = useState(false)

    document.title="喜欢的文章"


    useEffect(() => {
        if (articlesError) {
            message.error(articlesError + '')
        }
        if (articleList && articleList.CurLikeArticles.edges) {
            setList(articleList.CurLikeArticles.edges)
        } else {
            setList([])
        }
    }, [articleList, articlesError])

    const fetchData = () => {
        fetchMore({
            // @ts-ignore
            updateQuery: ({fetchMoreResult}: { fetchMoreResult: QueryResult }) => {
                const newEdges = fetchMoreResult.data.HotArticles.edges;
                const pageInfo = fetchMoreResult.data.HotArticles.pageInfo;
                setList(list.concat(...newEdges))
                setHasMore(pageInfo.hasNextPage)
            }
        },).catch((reason: any) => message.error(reason + ''))

    }

    return (
        <Row>
            <Col xs={24} sm={24} md={18} lg={18} xl={16} xxl={16}>
                {
                    // @ts-ignore
                    <UserInfo loading={false} data={{User: props.currentUser?.CurrentUser}}
                              currentUserId={props.currentUser?.CurrentUser.id}/>
                }
                <Tabs defaultActiveKey="article" tabBarStyle={{display: "flex", paddingLeft: 50, fontSize: 16}}>
                    <TabPane key='article' tab='喜欢的文章'>
                        <ArticleList
                            curId={props.currentUser?.CurrentUser.id}
                            fetchData={fetchData}
                            loading={articlesLoading && list}
                            hasMore={hasMore}
                            data={list}
                            locate={{emptyText: <Empty description="还没喜欢任何文章哦😦"/>}}
                        />
                    </TabPane>
                </Tabs>
            </Col>
            <Col xs={0} sm={0} md={6} lg={6} xl={8} xxl={8}>
                {
                    // @ts-ignore
                    <Introduce data={{User: props.currentUser?.CurrentUser}} currentUser={props.currentUser}/>
                }
            </Col>
        </Row>
    )
}