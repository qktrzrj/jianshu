import React, {useEffect, useState} from "react";
import {Link, RouteComponentProps} from "react-router-dom";
import {
    useReplyListQuery,
    CurrentUserQuery,
    ObjType,
    useAddCommentMutation, useAddReplyMutation,
    useArticleQuery,
    useHasLikeLazyQuery,
    useLikeMutation,
    useUnLikeMutation,
    useViewAddMutation, useAddMsgMutation, MsgType
} from "../../generated/graphql";
import {
    Affix,
    Avatar,
    Button,
    Card,
    Col,
    Comment,
    Divider,
    Form,
    Input,
    List,
    message,
    Row,
    Skeleton,
    Typography
} from "antd";
import {LikeOutlined} from '@ant-design/icons'
import ResultPage from "../../component/result/result";
import moment from 'moment'
import MarkdownIt from "markdown-it";
import hljs from "highlight.js";
import UserInfo from "../../component/userInfo/userInfo";
import {IconFont} from "../../component/IconFont";
import './index.less'
import 'highlight.js/styles/atom-one-light.css'

const {TextArea} = Input;


export default function Article(props: RouteComponentProps & { currentUser: CurrentUserQuery | undefined }) {

    const mdParser = new MarkdownIt({
        html: true,
        linkify: true,
        typographer: true,
        highlight(str, lang) {
            if (lang && hljs.getLanguage(lang)) {
                try {
                    return '<div class="highlight" style="padding: 10px"><div class="chroma">\n' +
                        '<table class="lntable"><tbody><tr><td class="lntd">\n' +
                        '<pre class="hljs" style="background-color: #f1f1f1"><code>' +
                        hljs.highlight(lang, str, true).value +
                        '</code></pre></td></tr></tbody></table>\n' +
                        '</div>\n' +
                        '</div>'
                } catch (__) {
                }
            }

            return '';
        },

    })

    const {data, error, loading} = useArticleQuery({
        // @ts-ignore
        variables: {id: parseInt(props.match.params.id)},
        fetchPolicy: "no-cache"
    })

    const [HasLike, {data: isLike, error: likeErr}] = useHasLikeLazyQuery()

    const [viewAdd] = useViewAddMutation()
    const [like] = useLikeMutation()
    const [unlike] = useUnLikeMutation()
    const [addComment] = useAddCommentMutation()
    const [addMsg] = useAddMsgMutation()

    const [comment, setComment] = useState('')
    const [likeNum, setLikeNum] = useState(0)
    const [hasLike, setHasLike] = useState(false)
    const [cmList, setCmList] = useState()

    useEffect(() => {
        if (data) {
            document.title = data.Article.title + ' - ' + data.Article.User.username
            setLikeNum(data.Article.LikeNum)
            setCmList(data.Article.CommentList)
            if (props.currentUser?.CurrentUser) {
                HasLike({variables: {id: data.Article.id, objtyp: ObjType.ArticleObj}})
            }
        } else {
            document.title = '指针'
        }
    }, [HasLike, data, props.currentUser])

    useEffect(() => {
        if (likeErr) {
            message.error(likeErr.message)
        }
        if (isLike) {
            setHasLike(isLike.HasLike)
        }
    }, [isLike, likeErr])

    useEffect(() => {
        viewAdd({
            // @ts-ignore
            variables: {id: parseInt(props.match.params.id)}
        }).catch(reason => message.error(reason + ''))
        // @ts-ignore
    }, [props.match.params.id, viewAdd])

    const Like = () => {
        if (data && props.currentUser?.CurrentUser) {
            if (hasLike) {
                unlike({variables: {id: data.Article.id, objtyp: ObjType.ArticleObj}})
                    .then(res => {
                        if (res.data && res.data.Unlike) {
                            setLikeNum(likeNum - 1)
                            setHasLike(false)
                        }
                    })
                    .catch(reason => message.error(reason + ''))
            } else {
                like({variables: {id: data.Article.id, objtyp: ObjType.ArticleObj}})
                    .then(res => {
                        if (res.data && res.data.Like) {
                            setLikeNum(likeNum + 1)
                            setHasLike(true)
                            if (props.currentUser){
                                addMsg({
                                    variables: {
                                        typ: MsgType.LikeMsg,
                                        fromId: props.currentUser.CurrentUser.id,
                                        toId: data.Article.User.id,
                                        content: `赞了你`,
                                    }
                                }).catch(reason => message.error(reason+''))
                            }
                        }
                    })
                    .catch(reason => message.error(reason + ''))
            }
        }
    }

    const AddComment = () => {
        if (props.currentUser && data) {
            if (comment === "") {
                message.warning("评论内容不能为空")
                return
            }
            addComment({variables: {id: data.Article.id, content: comment}})
                .then(res => {
                    if (res.data) {
                        if (cmList) {
                            setCmList([res.data.AddComment].concat(...cmList))
                        } else {
                            setCmList([res.data.AddComment])
                        }
                        setComment("")
                        message.success("发布成功")
                       if (props.currentUser){
                           addMsg({
                               variables: {
                                   typ: MsgType.CommentMsg,
                                   fromId: props.currentUser.CurrentUser.id,
                                   toId: data.Article.User.id,
                                   content: `评论了你的文章<a href="/p/${data.Article.id}">《${data.Article.title}》</a>`,
                               }
                           }).catch(reason => message.error(reason+''))
                       }
                    }
                    if (res.errors) {
                        message.error(res.errors + '')
                    }
                })
                .catch(reason => message.error(reason + ''))
        }
    }


    if (error) {
        return (
            <ResultPage status={"error"} title={"获取文章信息失败"} subTitle={"错误"}/>
        )
    }

    return (
        <Row gutter={2} className='article'>
            <Col span={2}>
                <Affix offsetTop={300}>
                    <div>
                        <Button type={(hasLike && 'primary') || undefined}
                                shape={"circle"}
                                icon={<LikeOutlined/>}
                                onClick={Like}/>
                        <br/>
                        {likeNum > 0 && <span>{likeNum}赞</span>}
                    </div>
                </Affix>
            </Col>
            <Col span={20}>
                <Skeleton loading={loading}>
                    {data &&
                    <div>
                        <Typography.Title style={{display: "flex"}}>
                            {data.Article.title}
                        </Typography.Title>
                        <Card bordered={false} style={{display: "flex"}}>
                            <Card.Meta avatar={<Avatar src={data.Article.User.avatar}/>}
                                       title={
                                           <Link style={{display: "flex"}}
                                                 to={'/u/' + data.Article.User.id}>
                                               {data.Article.User.username}
                                           </Link>}
                                       description={<div>
                                           {moment(data.Article.updatedAt).format('YYYY.MM.DD HH:SS:MM')}
                                           <span style={{marginLeft: 10}}>阅读{data.Article.ViewNum}</span>
                                       </div>}
                            />
                        </Card>
                        <div style={{textAlign: "left"}}
                             dangerouslySetInnerHTML={{__html: mdParser.render(data.Article.content)}}/>
                        <div style={{display: "flex"}}>
                            <Button type={(hasLike && 'primary') || undefined}
                                    shape={"circle"}
                                    icon={<LikeOutlined/>}
                                    onClick={Like}/>
                            {likeNum > 0 &&
                            <span style={{paddingTop: 5, paddingLeft: 5}}>{likeNum}人点赞</span>}
                        </div>
                        <Divider type={"horizontal"}/>
                        {
                            // @ts-ignore
                            <UserInfo loading={false} data={{User: data.Article.User}}
                                      currentUserId={props.currentUser?.CurrentUser.id}/>
                        }
                    </div>
                    }
                    {props.currentUser && props.currentUser.CurrentUser.id !== data?.Article.User.id &&
                    <Comment avatar={props.currentUser.CurrentUser.avatar}
                             content={
                                 <Form>
                                     <Form.Item>
                                         <TextArea autoSize={{maxRows: 4, minRows: 4}}
                                                   onChange={event => setComment(event.target.value)}
                                                   value={comment}/>
                                     </Form.Item>
                                     <Form.Item>
                                         <Button htmlType="submit" shape={"round"} loading={false}
                                                 onClick={AddComment}
                                                 type="primary">
                                             发布
                                         </Button>
                                     </Form.Item>
                                 </Form>
                             }/>
                    }
                    {cmList && <CommentList currentUser={props.currentUser?.CurrentUser} comments={cmList}/>}
                </Skeleton>
            </Col>
        </Row>
    )
}

function CommentList({comments, currentUser}: any) {

    return (
        <List
            dataSource={comments}
            locale={{emptyText: <div/>}}
            itemLayout="horizontal"
            renderItem={(item: any) => <CommentItem currentUser={currentUser} item={item}/>
            }
        />
    )
}

function CommentItem({item, currentUser}: any) {

    const {data, loading, refetch} = useReplyListQuery({variables: {id: item.id}, fetchPolicy: "no-cache"})
    const [fetchLike, {data: isLike}] = useHasLikeLazyQuery()

    const [like] = useLikeMutation()
    const [unlike] = useUnLikeMutation()
    const [addReply] = useAddReplyMutation()

    const [HasLike, setHasLike] = useState(false)
    const [reply, setReply] = useState(false)
    const [replyContent, setReplyContent] = useState('')
    const [replyList, setReplyList] = useState()

    useEffect(() => {
        refetch({id: item.id})
    }, [item, refetch])

    useEffect(() => {
        if (data) {
            setReplyList(data.ReplyList)
        }
    }, [data])


    useEffect(() => {
        if (currentUser) {
            fetchLike({variables: {id: item.id, objtyp: ObjType.CommentObj}})
        }
    }, [currentUser, fetchLike, item])

    useEffect(() => {
        if (isLike) {
            setHasLike(isLike.HasLike)
        }
    }, [isLike])

    const Like = () => {
        if (HasLike) {
            unlike({variables: {id: item.id, objtyp: ObjType.CommentObj}})
                .then(res => {
                    if (res.data && res.data.Unlike) {
                        setHasLike(false)
                    }
                })
                .catch(reason => message.error(reason + ''))
        } else {
            like({variables: {id: item.id, objtyp: ObjType.CommentObj}})
                .then(res => {
                    if (res.data && res.data.Like) {
                        setHasLike(true)
                    }
                })
                .catch(reason => message.error(reason + ''))
        }
    }

    const OnClickReply = () => {
        if (reply) {
            setReply(false)
        } else {
            setReply(true)
        }
    }

    const AddReply = (res: any) => {
        if (res.errors) {
            message.error(res.errors + '')
        }
        if (res.data) {
            if (replyList) {
                setReplyList(replyList.concat(res.data.AddReply))
            } else {
                setReplyList([res.data.AddReply])
            }
            setReplyContent('')
            setReply(false)
        }
    }

    const OnAddReply = () => {
        var spanEle = document.createElement('span')
        spanEle.appendChild(document.createTextNode(replyContent))
        addReply({variables: {id: item.id, content: spanEle.outerHTML}})
            .then(AddReply)
            .catch(reason => message.error(reason + ''))
    }

    const AddReplyToReply = ({id, username, content}: { id: string, username: string, content: string }) => {
        var spanEle = document.createElement('span')
        spanEle.style.marginLeft = '10px'
        spanEle.appendChild(document.createTextNode(content))
        addReply({
            variables: {
                id: item.id, content:
                    `<a href="/u/${id}">@${username}</a>` + spanEle.outerHTML
            }
        })
            .then(AddReply)
            .catch(reason => message.error(reason + ''))
    }

    return (
        <Comment
            content={<span style={{display: "flex"}}>{item.content}</span>}
            author={
                <div>
                    <Link style={{display: "flex"}} to={'/u/' + item.User.id}>{item.User.username}</Link>
                    <span>{item.floor}楼 {moment(item.updatedAt).format('YYYY.MM.DD HH:SS:MM')}</span>
                </div>}
            avatar={<Avatar src={item.User.avatar}/>}
            actions={[
                (!HasLike && <span onClick={Like}><IconFont type="icon-dianzan-copy"/> 赞</span>) ||
                <span style={{color: "orange"}} onClick={Like}><IconFont type="icon-dianzan-copy-copy"/> 1</span>,
                <span onClick={OnClickReply}><IconFont type="icon-huifu"/> 回复</span>
            ]}
        >
            {reply &&
            <Form>
                <Form.Item>
                    <TextArea autoSize={{maxRows: 2, minRows: 2}}
                              onChange={event => setReplyContent(event.target.value)}
                              value={replyContent}/>
                </Form.Item>
                <Form.Item>
                    <Button htmlType="submit" shape={"round"} loading={false}
                            onClick={OnAddReply}
                            type="primary">
                        发布
                    </Button>
                    <Button shape={"round"} style={{marginLeft: 5}}
                            onClick={() => setReply(false)}
                            type="default">
                        取消
                    </Button>
                </Form.Item>
            </Form>
            }
            {replyList && <ReplyList loading={loading} onAdd={AddReplyToReply} replies={replyList}/>}
        </Comment>
    )
}

function ReplyList({replies, loading, onAdd}: any) {
    return (
        <List
            loading={loading}
            dataSource={replies}
            locale={{emptyText: <div/>}}
            itemLayout="horizontal"
            renderItem={(item: any) => <ReplyItem onAdd={onAdd} item={item}/>
            }
        />
    )
}

function ReplyItem({item, onAdd}: any) {

    const [reply, setReply] = useState(false)
    const [replyContent, setReplyContent] = useState('')

    const OnClickReply = () => {
        if (reply) {
            setReply(false)
        } else {
            setReply(true)
        }
    }

    return (
        <Comment
            content={<div dangerouslySetInnerHTML={{__html: item.content}} style={{textAlign: "left"}}/>}
            author={
                <div>
                    <Link style={{display: "flex"}} to={'/u/' + item.User.id}>{item.User.username}</Link>
                    <span>{moment(item.updatedAt).format('YYYY.MM.DD HH:SS:MM')}</span>
                </div>}
            avatar={<Avatar src={item.User.avatar}/>}
            actions={[
                <span onClick={OnClickReply}><IconFont type="icon-huifu"/> 回复</span>
            ]}
        >
            {reply &&
            <Form>
                <Form.Item>
                    <TextArea autoSize={{maxRows: 2, minRows: 2}}
                              onChange={event => setReplyContent(event.target.value)}
                              value={replyContent}/>
                </Form.Item>
                <Form.Item>
                    <Button htmlType="submit" shape={"round"} loading={false}
                            onClick={() => {
                                onAdd({
                                    id: item.User.id,
                                    username: item.User.username,
                                    content: replyContent
                                })
                                setReply(false)
                            }}
                            type="primary">
                        发布
                    </Button>
                    <Button shape={"round"} style={{marginLeft: 5}}
                            onClick={() => setReply(false)}
                            type="default">
                        取消
                    </Button>
                </Form.Item>
            </Form>
            }
        </Comment>
    )
}

