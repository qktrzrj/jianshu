import React, {useEffect, useRef, useState} from "react";
import {Button, Divider, Drawer, Dropdown, Input, Layout, Menu, message, Spin, Steps, Typography, Upload} from "antd";
import './writer.less'
import MarkdownIt from 'markdown-it'
import MdEditor from 'react-markdown-editor-lite'
import 'react-markdown-editor-lite/lib/index.css'
import {Link} from "react-router-dom";
import {IconFont} from "../../component/IconFont";
import Publish from "../../component/mdplugins/publish";
import hljs from 'highlight.js'
import 'highlight.js/styles/atom-one-light.css'
import {
    ArticleEdge, ArticleState, UpdateArticleMutationVariables,
    useArticleLazyQuery,
    useDeleteArticleMutation,
    useDraftArticleMutation,
    useMyArticlesQuery, useNewArticleMutation, useUpdateArticleMutation, useUploadMutation
} from "../../generated/graphql";
import ReactDOM from "react-dom";
import {LoadingOutlined, PlusOutlined} from '@ant-design/icons';
import {RcCustomRequestOptions} from "antd/lib/upload/interface";


const {Step} = Steps;
const {Sider, Content} = Layout;
const {Text} = Typography;

/**
 * writer router component
 * @constructor
 */
export default function Writer() {
    const mdParser = new MarkdownIt({
        html: true,
        linkify: true,
        typographer: true,
        highlight(str, lang) {
            if (lang && hljs.getLanguage(lang)) {
                try {
                    return '<div class="highlight"><div class="chroma ">\n' +
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


    const {refetch: refetchMyArticles, data, loading, error} = useMyArticlesQuery()
    const [draft, {error: draftArticleError}] = useDraftArticleMutation()
    const [fetchArticle, {data: articleData, refetch, error: articleErr}] = useArticleLazyQuery()
    const [deleteArticle] = useDeleteArticleMutation()
    const [update, {loading: updateLoading}] = useUpdateArticleMutation()
    const [publish] = useNewArticleMutation()
    const [upload] = useUploadMutation()


    const [mdEditor, setMdEditor] = useState()
    const [drawer, setDrawer] = useState(false)
    const sk = useRef('')
    const [title, setTitle] = useState('')
    const [list, setList] = useState()
    const [content, setContent] = useState()
    const [introduce, setIntroduce] = useState('')
    const inputEnd = useRef(true)
    const articleState = useRef('发布文章')
    const [imageUrl, setImageUrl] = useState()


    const onPublish = () => {
        if (sk.current !== '') {
            publish({variables: {id: parseInt(sk.current)}})
                .then(res => {
                    if (res.errors) {
                        message.error("发布失败：" + res.errors)
                    }
                    if (res.data && res.data.NewArticle.id) {
                        message.success("发布成功")
                    }
                })
                .catch(reason => message.error(reason + ''))
        }
    }

    MdEditor.use(Publish, {
        OnClick: onPublish,
        State: articleState.current,
    })


    const updateArticle = (data: UpdateArticleMutationVariables, errMsg: string, refresh: boolean) => {
        update({variables: data}).then(res => {
            if (res.errors) {
                console.log(res.errors)
                message.error(errMsg)
            }
            if (res.data) {
                setTitle(res.data.UpdateArticle.title)
                switch (res.data.UpdateArticle.state) {
                    case ArticleState.Updated: {
                        articleState.current = '发布更新'
                        break
                    }
                    default: {
                        articleState.current = '发布文章'
                    }
                }
                if (refresh) {
                    updateListByRefetch()
                }
            }
        }).catch(reason => {
            message.error(reason + '')
        })
    }


    useEffect(() => {
        if (data && !list) {
            setList([])
            if (data.CurArticles.edges) {
                setList(data.CurArticles.edges)
                sk.current = data.CurArticles.edges[0].node.id + ''
                setTitle(data.CurArticles.edges[0].node.title)
                fetchArticle({variables: {id: data.CurArticles.edges[0].node.id}})
            }

        }
        if (articleErr && !content && content !== '') {
            message.error(articleErr)
        }
        if (articleData && !content && content !== '') {
            setContent(articleData.Article.content)
        }

    }, [articleData, articleErr, content, data, fetchArticle, list])


    useEffect(() => {
        if (mdEditor) {
            let mdInputs = document.getElementsByClassName('mdArea')
            let mdDOM = ReactDOM.findDOMNode(mdInputs[0]);
            mdDOM?.addEventListener('compositionstart', () => {
                inputEnd.current = false
            })
            mdDOM?.addEventListener('compositionend', () => {
                inputEnd.current = true
                console.log("结束输入")
            })
            mdDOM?.addEventListener('change', () => {
                if (inputEnd.current && sk.current !== '') {
                    updateArticle({
                        title: '',
                        id: parseInt(sk.current),
                        content: mdEditor.getMdValue(),
                        cover: null,
                        subTitle: null,
                    }, '保存失败', false)
                }
            })
        }
        // eslint-disable-next-line
    }, [mdEditor])

    const changeSelectContent = ({key}: { key: string }) => {
        sk.current = key
        refetch({id: parseInt(key)})
            .then(res => {
                if (res.errors) {
                    message.error(res.errors + '')
                }
                if (res.data) {
                    switch (res.data.Article.state) {
                        case ArticleState.Updated: {
                            articleState.current = '发布更新'
                            break
                        }
                        default: {
                            articleState.current = '发布文章'
                        }
                    }
                    setTitle(res.data.Article.title)
                    setContent(res.data.Article.content)
                }
            })
            .catch(reason => {
                message.error(reason + '')
            })
    }

    const updateListByRefetch = () => {
        refetchMyArticles().then(res => {
            if (res.errors) {
                message.error(res.errors + '')
            }
            if (res.data) {
                setList(res.data.CurArticles.edges)
            }
        }).catch(e => message.error(e + ''))
    }

    const updateMarkdown = ({text}: { html: string, text: string }) => {
        setContent(text)
    }

    const renderHTML = (text: string) => {
        return mdParser.render(text)
    }

    useEffect(() => {
        if (draftArticleError && draftArticleError.networkError && draftArticleError.networkError.message.includes("401")) {
            window.open('/signIn')
        }
    }, [draftArticleError])


    const customRequest = (options: RcCustomRequestOptions) => {
        upload({variables: {file: options.file}})
            .then(res => {
                if (res.errors) {
                    message.error(res.errors + '')
                }
                if (res.data) {
                    setImageUrl('http://localhost:8008/image/' + res.data.Upload)
                }
            })
            .catch(reason => message.error(reason + ''))
    }

    const onDelete = (id: number) => {
        deleteArticle({variables: {id: id}})
            .then(res => {
                if (res.errors) {
                    message.error(res.errors + '')
                }
                if (res.data) {
                    if (res.data.DeleteArticle) {
                        let last = -1;
                        list.forEach((node: any, i: number) => {
                            if (node.node.id === id) {
                                last = i - 1
                            }
                        })
                        const listcopy = list.filter((n: any) => n.node.id !== id);
                        if (listcopy.length) {
                            if (last >= 0) {
                                changeSelectContent({key: listcopy[last].node.id + ''})
                            } else {
                                changeSelectContent({key: listcopy[0].node.id + ''})
                            }
                        } else {
                            sk.current = ''
                            setTitle('')
                            setContent('')
                        }
                        setList(listcopy)
                        message.success("删除成功")
                    } else {
                        message.warning("删除失败")
                    }
                }
            })
            .catch(reason => message.error(reason + ''))
    }

    const onDraft = () => {
        let date = new Date();
        draft({
            variables: {
                title: date.getFullYear() + '-' + date.getMonth() + '-' + date.getDay(),
            }
        })
            .then(res => {
                if (res.errors) {
                    console.log(res.errors)
                    message.error('新建文章失败')
                }
                if (res.data) {
                    sk.current = res.data.DraftArticle.id + ''
                    articleState.current = '发布文章'
                    setTitle(res.data.DraftArticle.title)
                    setContent('')
                    setList([{node: res.data.DraftArticle}].concat(...list.slice()))
                }
            })
            .catch(reason => message.error(reason + ''))
    }

    const updateTitle = () => {
        if (!list || list.size === 0) {
            message.warning("当前没有可编辑的文章")
            return
        }
        updateArticle({
            title: title,
            id: parseInt(sk.current),
            content: null,
            cover: null,
            subTitle: null,
        }, '修改标题失败', true)
    }

    return (
        <Layout className="writer">
            <Sider className='sider'>
                <div className='index'>
                    <Link to='/'>
                        回首页
                    </Link>
                </div>
                <Divider type={"horizontal"}/>
                <div className='new' onClick={onDraft}>
                    <IconFont type='icon-jia' className='fa'/>新建文章
                </div>
                {error && <Text type='danger' style={{paddingLeft: 20}}>无法获取文章列表信息</Text>}
                <Spin spinning={loading}>
                    {list && list.length > 0 &&
                    <Menu mode={"vertical"} className='article-list-menu'
                          defaultSelectedKeys={[list.values().next().value.node.id + '']}
                          selectedKeys={[sk.current]}
                          onSelect={changeSelectContent}>
                        {list.map((item: ArticleEdge) => (
                            <Menu.Item key={item.node.id}>
                                <span>
                                    {item.node.title}
                                    {sk.current === item.node.id + '' && (
                                        <Dropdown overlay={
                                            <ArticleOpMenu
                                                delete={() => {
                                                    onDelete(item.node.id)
                                                }}
                                                setStyle={() => setDrawer(true)}
                                                publish={onPublish}
                                            />}
                                                  trigger={['click']}>
                                            <IconFont style={{float: 'right', paddingTop: 12}} type='icon-shezhi'/>
                                        </Dropdown>)
                                    }
                                </span>
                            </Menu.Item>
                        ))}
                    </Menu>
                    }
                </Spin>
            </Sider>
            <Layout>
                <Content>
                    <Drawer
                        title="设置发布样式"
                        onClose={() => setDrawer(false)}
                        visible={drawer}
                        width={720}
                        footer={
                            <div style={{textAlign: 'right'}}>
                                <Button onClick={() => setDrawer(false)} style={{marginRight: 8}}>
                                    取消
                                </Button>
                                <Button onClick={() => {
                                    updateArticle({
                                        id: parseInt(sk.current),
                                        content: null,
                                        title: null,
                                        cover: imageUrl,
                                        subTitle: introduce
                                    }, '设置发布样式失败', false)
                                    setDrawer(false)
                                }} type="primary">
                                    保存
                                </Button>
                            </div>
                        }
                    >
                        <Steps direction="vertical">
                            <Step title="上传封面图" status={"process"} description={
                                <div>
                                    小于10MB，格式 png/jpg
                                    <Upload name="cover" listType={"picture-card"} showUploadList={false}
                                            beforeUpload={beforeUpload}
                                            action="https://upload-z2.qiniup.com/"
                                            customRequest={customRequest}
                                    >
                                        {imageUrl ? <img src={imageUrl} alt="avatar" style={{width: '100%'}}/> : <div>
                                            {updateLoading ? <LoadingOutlined/> : <PlusOutlined/>}
                                            <div className="ant-upload-text">Upload</div>
                                        </div>}
                                    </Upload>
                                </div>
                            }/>
                            <Step title="输入摘要" status={"process"} description={
                                <Input.TextArea rows={3}
                                                value={introduce}
                                                onChange={(event) => setIntroduce(event.target.value)}
                                                autoSize={{maxRows: 3, minRows: 3}}
                                                placeholder="选填，如不填将抓取文章首段内容"/>
                            }/>
                        </Steps>
                    </Drawer>
                    <Input value={title}
                           onChange={event => {
                               setTitle(event.target.value)
                           }}
                           suffix={<Button onClick={updateTitle}>修改</Button>}/>
                    <MdEditor
                        ref={node => {
                            if (node !== null) {
                                setMdEditor(node)
                            }
                        }}
                        value={content}
                        style={{height: document.body.clientHeight - 42, width: '100%'}}
                        config={{
                            view: {
                                menu: true,
                                md: true,
                                html: true,
                                fullScreen: true,
                                hideMenu: false,
                            },
                            imageAccept: '.jpg,.png',
                            markdownClass: 'mdArea',
                            htmlClass: 'htmlArea',
                            table: {
                                maxRow: 5,
                                maxCol: 6,
                            },
                            syncScrollMode: ['leftFollowRight', 'rightFollowLeft'],
                        }}
                        renderHTML={renderHTML}
                        onChange={updateMarkdown}
                    />
                </Content>
            </Layout>
        </Layout>
    )
}

function ArticleOpMenu(props: { delete: () => void, setStyle: () => void, publish: () => void }) {

    return (
        <Menu>
            <Menu.Item onClick={props.publish}>直接发布</Menu.Item>
            <Menu.Item onClick={props.setStyle}>设置发布样式</Menu.Item>
            <Menu.Item onClick={props.delete}>删除文章</Menu.Item>
        </Menu>
    )
}

function beforeUpload(file: any) {
    const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png';
    if (!isJpgOrPng) {
        message.error('你只能上传格式为 JPG/PNG 的文件!');
    }
    const isLt10M = file.size / 1024 / 1024 < 10;
    if (!isLt10M) {
        message.error('文件大小必须小于 10MB!');
    }
    return isJpgOrPng && isLt10M;
}