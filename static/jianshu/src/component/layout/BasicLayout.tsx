import {Link, RouteComponentProps} from "react-router-dom";
import React, {ReactNode, useEffect, useState} from "react";
import BasicLayout from "@ant-design/pro-layout";
import {Route} from "@ant-design/pro-layout/es/typings";
import Right from "./right/right";
import {BackTop, Badge, Layout as AntLayout} from "antd";
import "./basicLayout.less"
import {CurrentUserQuery, useCurrentUserQuery} from "../../generated/graphql";
import {useLazyQuery} from "@apollo/react-hooks";
import {MsgNumGQl} from "../query/query";

const {Content} = AntLayout;

const menuData1: Route = {
    routes: [
        {
            icon: "icon-faxianx",
            name: "首页",
            key: "index",
            path: "/"
        },
    ]
}

const menuData2: Route = {
    routes: [
        {
            icon: "icon-faxianx",
            name: "发现",
            key: "found",
            path: "/"
        },
        {
            icon: "icon-2",
            name: "关注",
            key: "subscriptions",
            path: "/subscriptions"
        },
        {
            icon: "icon-lingdang",
            name: "消息",
            key: "notifications",
            path: "/notifications",
            routes: [
                {
                    icon: "icon-pinglun",
                    name: "评论",
                    key: "comments",
                    path: "/notifications/comments"
                },
                {
                    icon: "icon-xihuan",
                    name: "喜欢",
                    key: "likes",
                    path: "/notifications/likes"
                },
                {
                    icon: "icon-guanzhu2",
                    name: "关注",
                    key: "follows",
                    path: "/notifications/follows"
                },
                {
                    icon: "icon-iconqita",
                    name: "回复",
                    key: "replies",
                    path: "/notifications/replies"
                },
            ]
        },
    ]
}

function Layout(props: RouteComponentProps & { render: (props: CurrentUserQuery | undefined) => React.ReactNode }) {
    const {loading, data} = useCurrentUserQuery()

    const [fetch, {data: Num}] = useLazyQuery(MsgNumGQl, {pollInterval: 5000})

    const [totNum, setTotNum] = useState(0)
    const [cmtNum, setCmtNum] = useState(0)
    const [replyNum, setReplyNUm] = useState(0)
    const [likeNum, setLikeNum] = useState(0)
    const [followNum, setFollowNum] = useState(0)

    useEffect(() => {
        if (data) {
            fetch()
        }
    }, [fetch, data])

    useEffect(() => {
        if (Num) {
            setTotNum(Num.MsgNum.comment + Num.MsgNum.like + Num.MsgNum.follow + Num.MsgNum.reply)
            setCmtNum(Num.MsgNum.comment)
            setLikeNum(Num.MsgNum.like)
            setFollowNum(Num.MsgNum.follow)
            setReplyNUm(Num.MsgNum.reply)
        }
    }, [Num])

    return (
        <BasicLayout
            logo="/logo192.png"
            title=""
            menuHeaderRender={(logoDom, titleDom) => (
                <Link to="/" target='_top'>
                    {logoDom}
                    {titleDom}
                </Link>
            )}
            layout={"topmenu"}
            contentWidth="Fixed"
            navTheme="light"
            fixedHeader
            iconfontUrl="//at.alicdn.com/t/font_1550295_4bnx9xx025.js"
            pageTitleRender={({breadcrumbMap, pathname}) => {
                let r = breadcrumbMap ? breadcrumbMap.get(pathname ? pathname : "") : {}
                return r ? r.name ? r.name + "" : "" : ""
            }}
            subMenuItemRender={(item, dom) => {
                return <Badge offset={[10, 0]} count={totNum}>{dom}</Badge>
            }}
            menuItemRender={(menuItemProps, defaultDom) => {
                if (menuItemProps.isUrl || menuItemProps.children || !menuItemProps.path) {
                    return defaultDom;
                }
                switch (menuItemProps.name) {
                    case '评论': {
                        return <Badge offset={[10, 0]} count={cmtNum}>
                            <Link to={menuItemProps.path}>{defaultDom}</Link></Badge>
                    }
                    case '喜欢': {
                        return <Badge offset={[10, 0]} count={likeNum}>
                            <Link to={menuItemProps.path}>{defaultDom}</Link></Badge>
                    }
                    case '关注': {
                        return <Badge offset={[10, 0]} count={followNum}>
                            <Link to={menuItemProps.path}>{defaultDom}</Link></Badge>
                    }
                    case '回复': {
                        return <Badge offset={[10, 0]} count={replyNum}>
                            <Link to={menuItemProps.path}>{defaultDom}</Link></Badge>
                    }
                }
                return <Link to={menuItemProps.path}>{defaultDom}</Link>;
            }}
            route={(data?.CurrentUser.id !== undefined && menuData2) || menuData1}
            rightContentRender={() => <Right data={data} Route={props} loading={loading}/>}
            contentStyle={{backgroundColor: "#fff", margin: "0"}}
        >
            <AntLayout className="content">
                <Content style={{paddingTop: 30}}>
                    {data && props.render({...data})}
                    {!data && props.render(data)}
                    <BackTop/>
                </Content>
            </AntLayout>
        </BasicLayout>
    )
}

export default Layout;