import React, { useEffect } from 'react';
import NProgress from 'nprogress';
import 'nprogress/nprogress.css';
import {Col, Row} from "antd";

const Loading = () => {
    useEffect(() => {
        NProgress.start();
        return () => {
            NProgress.done();
        };
    }, []);
    return (
        <Row>
            <Col span={12} offset={6}>

            </Col>
        </Row>
    );
};

export default Loading;