import {Button, Result,Typography} from "antd";
import React from "react";
import {CloseCircleOutlined} from '@ant-design/icons';
import {ResultStatusType} from "antd/lib/result";


const {Paragraph,Text} = Typography;

export default function ResultPage(props:{status:ResultStatusType,title:string,subTitle:string}) {
    return (
        <Result
            status={props.status}
            title={props.title}
            subTitle={props.subTitle}
            extra={
                [
                    <Button type="primary" key="refresh" onClick={() => window.location.reload()}>
                        刷新
                    </Button>,
                    <Button key="call">联系管理员</Button>,
                ]
            }>
            <div className="desc">
                <Paragraph>
                    <Text
                        strong
                        style={{
                            fontSize: 16,
                        }}
                    >
                        造成网页错误的原因可能是:
                    </Text>
                </Paragraph>
                <Paragraph>
                    <CloseCircleOutlined className="site-result-error-icon"/> 网络错误
                </Paragraph>
                <Paragraph>
                    <CloseCircleOutlined className="site-result-error-icon"/> 网站服务崩溃
                </Paragraph>
            </div>
        </Result>)
}