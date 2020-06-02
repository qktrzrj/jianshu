import React, {useEffect, useState} from "react";
import {Button, Card, Col, Empty, Form, Input, message, Row, Tabs} from "antd";
import {RouteComponentProps} from "react-router-dom";
import {
    CurrentUserQuery,
    useArticlesLazyQuery,
    useUpdateUserInfoMutation,
    useUserQuery
} from "../../generated/graphql";
import ResultPage from "../../component/result/result";
import './userCenter.less'
import UserInfo from "../../component/userInfo/userInfo";
import ArticleList from "../../component/articleList/articleList";
import {QueryResult} from "@apollo/react-common";
import {IconFont} from "../../component/IconFont";

const {TabPane} = Tabs

export default function UserCenter(props: RouteComponentProps & { currentUser: CurrentUserQuery | undefined }) {

    // @ts-ignore
    const {data, loading, error} = useUserQuery({variables: {id: parseInt(props.match.params.id)}})
    const [articles, {data: articleList, error: articlesError, loading: articlesLoading, fetchMore}] = useArticlesLazyQuery()
    const [update] = useUpdateUserInfoMutation()

    const [list, setList] = useState()
    const [hasMore, setHasMore] = useState(false)
    const [edit, setEdit] = useState(false)
    const [introduce, setIntroduce] = useState('')

    useEffect(() => {
        if (data) {
            document.title = data.User.username
            setList([])
            setIntroduce(data.User.introduce)
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

    const onFinish = () => {
        update({variables: {username:null,avatar: null, email: null, password: null, gender: null, introduce: introduce}})
            .then(res => {
                if (res.errors) {
                    setIntroduce(data?.User.introduce)
                    message.error(res.errors + '')
                }
                if (res.data) {
                    if (!res.data.UpdateUserInfo) {
                        setIntroduce(data?.User.introduce)
                    }
                }
            })
            .catch(reason => {
                setIntroduce(data?.User.introduce)
                message.error(reason + '')
            })
        setEdit(false)
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
                <Card bordered={false} className='introduce-form' title={<span style={{float: "left"}}>个人介绍</span>}
                      extra={data?.User.id === props.currentUser?.CurrentUser?.id &&
                      // eslint-disable-next-line no-script-url,jsx-a11y/anchor-is-valid
                      [<a href={'javascript:void(0)'} onClick={() => setEdit(true)}><IconFont type='icon-xie'/>编辑</a>]}>
                    {!edit && introduce}
                    {edit && <Form onFinish={onFinish}>
                        <Form.Item>
                            <Input.TextArea value={introduce} className='text-a'
                                            autoSize={{minRows: 5, maxRows: 5}}
                                            onChange={e => setIntroduce(e.target.value)}
                            />
                        </Form.Item>
                        <Form.Item style={{float: "left"}}>
                            <Button htmlType='submit' className='btn-hollow'>保存</Button>
                            {/* eslint-disable-next-line no-script-url,jsx-a11y/anchor-is-valid */}
                            <a href={'javascript:void(0)'} className='btn-cancel'
                               onClick={() => {
                                   setEdit(false)
                                   setIntroduce(data?.User.introduce)
                               }}>
                                取消
                            </a>
                        </Form.Item>
                    </Form>}
                </Card>
            </Col>
        </Row>
    )
}