import React, {useEffect, useState} from "react";
import "./index.less"
import {Empty, message} from "antd";
import {QueryResult} from '@apollo/react-common';
import {useHotArticlesQuery} from "../../generated/graphql";
import ResultPage from "../../component/result/result";
import ArticleList from "../../component/articleList/articleList";


/**
 * Index router component
 * @constructor
 */
export default function Index() {

    const {data, loading, error, fetchMore} = useHotArticlesQuery({variables: {cursor: null}})

    const [list, setList] = useState()
    const [hasMore, setHasMore] = useState(false)

    useEffect(() => {
        if (data) {
            setList(data.HotArticles.edges || [])
        }
    }, [data])


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

    if (error) {
        return (
            <ResultPage status={"error"} title={"奔溃通知书"} subTitle={"首页崩溃了😭！暂时连接不上..."}/>
        )
    }

    return (
        <ArticleList
            curId={undefined}
            fetchData={fetchData}
            loading={loading}
            hasMore={hasMore}
            data={list}
            locate={{emptyText: <Empty description="还没人发表文章哦😦"/>}}
        />
    )
}