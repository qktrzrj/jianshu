import {RouteComponentProps} from "react-router-dom";
import {Button, Result} from "antd";
import React from "react";

export default function noMatch(props: RouteComponentProps) {

    document.title = '404'

    return (
        <Result
            status={"warning"}
            title={"404"}
            subTitle={"页面不见了😭！"}
            extra={
                [
                    <Button type="primary" key="index" onClick={() => props.history.push('/')}>
                        回首页
                    </Button>,
                    <Button key="refresh" onClick={() => props.history.goBack()}>
                        返回上一页
                    </Button>,
                ]
            }>
        </Result>
    )
}