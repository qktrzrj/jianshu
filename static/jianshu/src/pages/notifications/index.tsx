import React, {useEffect, useState} from "react";
import {Avatar, Card, List, Skeleton, Tabs} from "antd";
import {Link, RouteComponentProps} from "react-router-dom";
import {CurrentUserQuery, MsgType, useListMsgLazyQuery} from "../../generated/graphql";
import moment from 'moment'
import ResultPage from "../../component/result/result";
import {IconFont} from "../../component/IconFont";

export default function Notifications(props: RouteComponentProps & { currentUser: CurrentUserQuery }) {

    const [fetch, {data, error, loading}] = useListMsgLazyQuery({fetchPolicy: "no-cache"})

    //@ts-ignore
    const [key, setKey] = useState(props.match.params.key)

    const getKey = (key: string) => {
        switch (key) {
            case 'comments': {
                return MsgType.CommentMsg
            }
            case 'likes': {
                return MsgType.LikeMsg
            }
            case 'follows': {
                return MsgType.FollowMsg
            }
            case 'replies': {
                return MsgType.ReplyMsg
            }
        }
        return MsgType.CommentMsg
    }

    useEffect(() => {
        //@ts-ignore
        setKey(props.match.params.key)
        //@ts-ignore
    }, [props.match.params.key])

    useEffect(() => {
        fetch({variables: {typ: getKey(key)}})
    }, [fetch, key])

    const onChangeKey = (key: string) => {
        props.history.push('/notifications/' + key)
    }

    const CardList = () => <Skeleton loading={loading}>
        <List dataSource={data?.ListMsg || []}
              locale={{emptyText: <div/>}}
              renderItem={(item) => {
                  return <Card bordered={false}>
                      <Card.Meta avatar={<Avatar src={item.User.avatar}/>}
                                 description={<div style={{textAlign: "left"}}>
                                     <Link
                                         to={'/u/' + item.User.id}>{item.User.username}</Link>
                                     <span style={{marginLeft: 10}}
                                           dangerouslySetInnerHTML={{__html: item.content}}/>
                                     <br/>
                                     <span>{moment(item.updatedAt).format('YYYY.MM.DD HH:SS:MM')}</span>
                                 </div>}
                      />
                  </Card>
              }}
        />
    </Skeleton>


    if (error) {
        return <ResultPage status={"error"} title={"错误"} subTitle={"获取消息错误"}/>
    }


    return (
        <Tabs
              activeKey={key}
              tabPosition={"left"}
              onChange={onChangeKey}
        >
            <Tabs.TabPane key='comments' tab={<span><IconFont type="icon-pinglun"/>评论</span>}>
                <CardList/>
            </Tabs.TabPane>
            <Tabs.TabPane key='likes' tab={<span><IconFont type="icon-xihuan"/>喜欢</span>}>
                <CardList/>
            </Tabs.TabPane>
            <Tabs.TabPane key='follows' tab={<span><IconFont type="icon-guanzhu2"/>关注</span>}>
                <CardList/>
            </Tabs.TabPane>
            <Tabs.TabPane key='replies' tab={<span><IconFont type="icon-iconqita"/>回复</span>}>
                <CardList/>
            </Tabs.TabPane>
        </Tabs>
    )
}