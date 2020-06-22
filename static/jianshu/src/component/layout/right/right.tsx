import {Menu, Row, Input, Dropdown, Avatar, Col, message, Spin} from "antd";
import React, {Suspense, useState} from "react";
import {Link, RouteComponentProps} from "react-router-dom";
import {IconFont} from "../../IconFont";
import "./right.less"
import Loading from "../../loading";
import {ClickParam} from "antd/es/menu";
import NProgress from 'nprogress';
import 'nprogress/nprogress.css';
import {useLogoutMutation} from "../../../generated/graphql";

const {Search} = Input;

interface right {
    data: any
    loading: boolean,
    Route: RouteComponentProps
}

export default function Right(props: right) {

    const [logout] = useLogoutMutation()

    const [Len, setLen] = useState(240)
    const [sval,setSval]=useState("")

    const onclick = (param: ClickParam) => {
        NProgress.start()
        switch (param.key) {
            case "index": {
                props.Route.history.push('/u/' + props.data.CurrentUser.id)
                break
            }
            case "like": {
                props.Route.history.push('/user/' + props.data.CurrentUser.id + '/liked_notes')
                break
            }
            case 'setting': {
                props.Route.history.push('/setting/basic')
                break
            }
            case "logout": {
                logout().then(r => {
                    if (r.errors) {
                        let err = r.errors.join("\n");
                        message.error(err);
                        return
                    }
                    if (r.data) {
                        window.location.href = '/'
                    }
                }).catch(reason => {
                    message.error(reason.toString());
                })
                break
            }
        }
        NProgress.done()
    }

    return (
        <Spin spinning={props.loading}>
            <Row style={{height: "64px", flexFlow: "nowrap"}} className="right" justify="end">

                <Col>
                    <Search style={{width: Len}}
                            onBlur={() => setLen(240)}
                            onFocus={() => setLen(400)}
                            placeholder="搜索"
                            size={"large"}
                            value={sval}
                            onChange={e=>setSval(e.target.value)}
                            onSearch={(value) => {
                                props.Route.history.replace('/search', {q: sval})
                                setSval("")
                            }}
                            className="search"/>
                </Col>

                {props.data && props.data.CurrentUser && props.data.CurrentUser.id &&
                <Col className="user">
                    <Dropdown className="avatar" overlay={
                        <Menu className="collapse" onClick={onclick}>
                            <Menu.Item key="index" className="item">
                                <IconFont type="icon-gerenzhongxin" className="item-icon"/>个人主页
                            </Menu.Item>
                            <Menu.Item key="like" className="item">
                                <IconFont type="icon-xihuan1" className="item-icon"/>喜欢的文章
                            </Menu.Item>
                            <Menu.Item key="setting" className="item">
                                <IconFont type="icon-shezhi" className="item-icon"/>设置
                            </Menu.Item>
                            <Menu.Item key="logout" className="item">
                                <IconFont type="icon-signout" className="item-icon"/>退出
                            </Menu.Item>
                        </Menu>
                    }>
                        <Avatar style={{top: 10}} size={"large"} src={props.data.CurrentUser.avatar}/>
                    </Dropdown>
                    <a className="write-btn" href='/writer' target='_blank'>
                        <IconFont type="icon-zuozhe"/>写文章
                    </a>
                </Col>
                }

                {(!props.data || !props.data.CurrentUser || !props.data.CurrentUser.id) &&
                <Col className="user">
                    <Suspense fallback={<Loading/>}>
                        <Link className="log-in-btn" to="/signIn">登录</Link>
                        <Link className="sign-up-btn" to="/signUp">注册</Link>
                    </Suspense>
                </Col>
                }
            </Row>
        </Spin>
    )
}