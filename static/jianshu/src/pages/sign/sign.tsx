import React, {useEffect, useState} from "react";
import "./sign.less"
import {Button, Checkbox, Form, Input, Layout, Tabs, Tooltip, message} from "antd";
import {RouteComponentProps} from "react-router-dom";
import {UserOutlined, LockOutlined, MailOutlined} from '@ant-design/icons';
import {useLazyQuery, useMutation} from "@apollo/react-hooks";
import {CheckEmailGQL, CheckUsernameGQL, SignInGQL, SignUpGQL} from '../../component/query/query'
import {IconFont} from "../../component/IconFont";
import NProgress from 'nprogress';
import 'nprogress/nprogress.css';

const {Header, Content} = Layout;
const {TabPane} = Tabs;

/**
 * sign router component
 * @constructor
 */
export default function Sign(props: RouteComponentProps) {
    const [path, setPath] = useState(props.location.pathname);
    useEffect(() => {
        setPath(props.location.pathname)
    }, [props.location.pathname]);
    return (
        <Layout className="sign">
            <Header className="logo">
                <a href="/">
                    <img style={{width: 100, height: 100}} src="logo192.png" alt="logo"/>
                </a>
            </Header>
            <Content className="main">
                <Tabs
                    tabBarGutter={60} defaultActiveKey={path} onTabClick={() => {
                    if (path === "/signIn") {
                        props.history.push("/signUp")
                    } else {
                        props.history.push("/signIn")
                    }
                }}>
                    <TabPane tab="登录" key="/signIn" className="title">
                        <LoginForm history={props.history} location={props.location} match={props.match}/>
                    </TabPane>
                    <TabPane tab="注册" key="/signUp" className="title">
                        <RegisterForm history={props.history} location={props.location} match={props.match}/>
                    </TabPane>
                </Tabs>
            </Content>
        </Layout>
    )
};

const LoginForm = (props: RouteComponentProps) => {

    const [form] = Form.useForm();

    const [signIn] = useMutation(SignInGQL);

    const onFinish = (values: any) => {
        if (values.rememberme===undefined){
            values.rememberme = false
        }
        NProgress.start()
        signIn({variables: values}).then(r => {
            if (r.errors) {
                let err = r.errors.join("\n");
                message.error(err);
                return
            }
            if (r.data) {
                props.history.push('/')
            }
        }).catch(reason => {
            message.error(reason.toString());
        })
        NProgress.done()
    };

    return (
        <Form form={form} style={{paddingTop: 10}} name="login" initialValues={{remember: true}} scrollToFirstError
              onFinish={onFinish}>
            <Form.Item name="username" className="input-prepend restyle" rules={
                [{required: true, message: '用户名不能为空'}]}>
                <Input size={"large"} prefix={<UserOutlined className="icon"/>}
                       placeholder="用户名 / 邮箱"/>
            </Form.Item>
            <Form.Item name="password" className="input-prepend" rules={
                [{required: true, message: '密码不能为空'}]
            }>
                <Input.Password prefix={<LockOutlined/>} size={"large"} placeholder="Password"/>
            </Form.Item>
            <Form.Item>
                <Form.Item name="rememberme" valuePropName="checked" className="remember-btn">
                    <Checkbox>记住我</Checkbox>
                </Form.Item>
                <div className="forgot-btn">
                    <a href="/?">
                        忘记密码
                    </a>
                </div>
            </Form.Item>
            <Form.Item>
                <Button htmlType="submit" className="login-btn">登录</Button>
            </Form.Item>
        </Form>
    )
};

const RegisterForm = (props: RouteComponentProps) => {

    const [form] = Form.useForm();

    const [uv, setuv] = useState(false);
    const [ev, setev] = useState(false);

    const [signUp] = useMutation(SignUpGQL);

    const onFinish = (values: any) => {
        NProgress.start()
        signUp({variables: values}).then(r => {

            if (r.errors) {
                let err = r.errors.join("\n");
                message.error(err);
                return
            }
            if (r.data) {
                props.history.push('/signIn')
            }
        }).catch(reason => {
            message.error(reason.toString());
        })
        NProgress.done()
    };

    const [checkUsername, {error: error1}] = useLazyQuery(CheckUsernameGQL);

    const [checkEmail, {error: error2}] = useLazyQuery(CheckEmailGQL);

    const validateUsername = () => {
        let username = form.getFieldValue("username");
        if (username) {
            checkUsername({variables: {username: username}});
            if (error1 && error1.message !== "") {
                setuv(true);
                return
            }
        }
        setuv(false);

    };

    const validateEmail = () => {
        let email = form.getFieldValue("email");
        if (email) {
            checkEmail({variables: {email: email}});
            if (error2 && error2.message !== "") {
                setev(true);
                return
            }
        }
        setev(false);
    };

    return (
        <Form form={form} style={{paddingTop: 10}} name="register" scrollToFirstError onFinish={onFinish}>

            <Form.Item name="username" className="input-prepend restyle" rules={[
                {max: 20, message: '用户名长度不能多于20个字符!'},
                {min: 6, message: '用户名长度最低为8个字符!'},
                {required: true, message: '请输入用户名'},
            ]}>
                <Input onBlur={validateUsername}
                       onFocus={() => setuv(false)}
                       size={"large"} prefix={<UserOutlined/>}
                       placeholder="设置你的用户名"
                       suffix={
                           <Tooltip visible={uv} getPopupContainer={() => document.body}
                                    placement={"right"} className="tooltip-inner"
                                    overlay={error1 && error1.message !== "" &&
                                    <div>
                                        <IconFont type="icon-gantanhao"/><span>  {error1.message}</span>
                                    </div>}
                           />}
                />
            </Form.Item>
            <Form.Item name="email" className="input-prepend restyle" rules={[
                {type: 'email', message: "请输入正确的邮箱地址"},
                {required: true, message: '请输入你的邮箱地址'},
            ]}>

                <Input onBlur={validateEmail}
                       onFocus={() => setev(false)}
                       size={"large"} prefix={<MailOutlined/>}
                       placeholder="设置邮箱地址"
                       suffix={
                           <Tooltip visible={ev} getPopupContainer={() => document.body} placement={"right"}
                                    overlay={error2 && error2.message !== "" &&
                                    <div className="tooltip-inner">
                                        <IconFont type="icon-gantanhao"/><span>  {error2.message}</span>
                                    </div>}
                           />}
                />
            </Form.Item>
            <Form.Item name="password" className="input-prepend" rules={[
                {max: 20, message: '密码长度不能多于20个字符'},
                {min: 8, message: '密码长度不能少于8个字符'},
                {required: true, message: '请设置你的密码'}
            ]}>
                <Input.Password prefix={<LockOutlined/>} size={"large"}
                                placeholder="设置密码"/>
            </Form.Item>
            <Form.Item>
                <Button htmlType="submit" className="register-btn">注册</Button>
            </Form.Item>
        </Form>
    )
};
