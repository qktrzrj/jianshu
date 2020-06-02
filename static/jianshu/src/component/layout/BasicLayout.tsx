import {Link, RouteComponentProps} from "react-router-dom";
import React from "react";
import BasicLayout from "@ant-design/pro-layout";
import {Route} from "@ant-design/pro-layout/es/typings";
import Right from "./right/right";
import {Layout as AntLayout} from "antd";
import "./basicLayout.less"
import {CurrentUserQuery, useCurrentUserQuery} from "../../generated/graphql";

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
                    name: "其他",
                    key: "others",
                    path: "/notifications/others"
                },
            ]
        },
    ]
}


function Layout(props: RouteComponentProps & { render: (props: CurrentUserQuery | undefined) => React.ReactNode }) {
    const {loading, data} = useCurrentUserQuery()

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
            iconfontUrl="//at.alicdn.com/t/font_1550295_8o4gpsdcwze.js"
            pageTitleRender={({breadcrumbMap, pathname}) => {
                let r = breadcrumbMap ? breadcrumbMap.get(pathname ? pathname : "") : {}
                return r ? r.name ? r.name + "" : "" : ""
            }}
            menuItemRender={(menuItemProps, defaultDom) => {
                if (menuItemProps.isUrl || menuItemProps.children || !menuItemProps.path) {
                    return defaultDom;
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
                </Content>
            </AntLayout>
        </BasicLayout>
    )
}

export default Layout;