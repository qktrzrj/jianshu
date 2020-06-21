import React, {useEffect, useState} from "react";
import {Avatar, Button, Card, message} from "antd";
import {Link} from "react-router-dom";
import {
    Gender, MsgType,
    useAddMsgMutation,
    useFollowMutation,
    useIsFollowLazyQuery,
    UserQuery,
    useUnFollowMutation
} from "../../generated/graphql";
import './userInfo.less'
import {IconFont} from "../IconFont";

export default function UserInfo(props: {
    loading: boolean,
    data: UserQuery | undefined,
    currentUserId: number | undefined,
    className?: string
}) {
    const {data, loading, currentUserId} = props

    const [fetch, {data: isFollow, error}] = useIsFollowLazyQuery()
    const [follow] = useFollowMutation()
    const [unFollow] = useUnFollowMutation()
    const [addMsg] = useAddMsgMutation()

    const [IsFollow, setIsFollow] = useState(false)

    useEffect(() => {
        if (currentUserId && data) {
            fetch({variables: {id: data.User.id}})
        }
    }, [currentUserId, data, fetch])

    useEffect(() => {
        if (isFollow) {
            setIsFollow(isFollow.IsFollow)
        }
        if (error) {
            message.error(error.message)
        }
    }, [error, isFollow])

    const Follow = () => {
        if (!currentUserId) {
            window.open('/signIn')
        } else if (data) {
            follow({variables: {id: data.User.id}})
                .then(res => {
                    if (res.errors) {
                        message.error(res.errors + '')
                    }
                    if (res.data && res.data.Follow) {
                        setIsFollow(true)
                        if (props.currentUserId){
                            addMsg({
                                variables: {
                                    typ: MsgType.FollowMsg,
                                    fromId: props.currentUserId,
                                    toId: data.User.id,
                                    content: `关注了你`,
                                }
                            }).catch(reason => message.error(reason+''))
                        }
                    }
                })
                .catch(e => message.error(e + ''))
        }
    }

    const UnFollow = () => {
        if (data) {
            unFollow({variables: {id: data.User.id}})
                .then(res => {
                    if (res.errors) {
                        message.error(res.errors + '')
                    }
                    if (res.data && res.data.UnFollow) {
                        setIsFollow(false)
                    }
                })
                .catch(e => message.error(e + ''))
        }
    }


    return (<Card
            bordered={false}
            loading={loading}
            size={"small"}
            className={props.className || 'user-card'}
        >
            <Card.Meta
                avatar={<Avatar className='avatar' size={"large"} src={data?.User.avatar}/>}
                title={<Link to={'/u/' + data?.User.id} className='name'>
                    {data?.User.username}
                    {data?.User.gender === Gender.Man && <IconFont className='gender' type='icon-male2'/>}
                    {data?.User.gender === Gender.Woman &&
                    <IconFont className='gender' type='icon-xingtaiduICON_nvxing'/>}
                    {data?.User.gender === Gender.Unknown && <IconFont className='gender' type='icon-privacy'/>}
                </Link>}
                description={
                    <Card bordered={false} className='info'>
                        <Card.Grid hoverable={false} className='meta'>
                            <span>{data?.User.FollowNum}</span>
                            <span>关注</span>
                        </Card.Grid>
                        <Card.Grid hoverable={false} className='meta'>
                            <span>{data?.User.FansNum}</span>
                            <span>粉丝</span>
                        </Card.Grid>
                        <Card.Grid hoverable={false} className='meta'>
                            <span>{data?.User.Words}</span>
                            <span>字数</span>
                        </Card.Grid>
                        <Card.Grid hoverable={false} className='meta'>
                            <span>{data?.User.LikeNum}</span>
                            <span>收获喜欢</span>
                        </Card.Grid>
                        {currentUserId !== data?.User.id && !IsFollow &&
                        <Button className="user-follow-btn" size={"large"} onClick={Follow}>+ 关注</Button>}
                        {currentUserId !== data?.User.id && IsFollow &&
                        <Button className="cancel-follow-btn" size={"large"} onClick={UnFollow}>√ 已关注</Button>}
                    </Card>
                }
            />
        </Card>
    )
}