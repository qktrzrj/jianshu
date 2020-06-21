import React, {useState} from "react";
import {Button, Card, Form, Input, message} from "antd";
import {IconFont} from "../IconFont";
import {CurrentUserQuery, UserQuery, useUpdateUserInfoMutation} from "../../generated/graphql";
import './introduce.less'

export default function Introduce(props: {
    data: UserQuery | undefined,
    currentUser: CurrentUserQuery | undefined,
}) {
    const {data} = props

    const [edit, setEdit] = useState(false)
    const [introduce, setIntroduce] = useState(data?.User.introduce)

    const [update] = useUpdateUserInfoMutation()


    const onFinish = () => {
        update({
            variables: {
                username: null,
                avatar: null,
                email: null,
                password: null,
                gender: null,
                introduce: introduce
            }
        })
            .then(res => {
                if (res.errors) {
                    setIntroduce(data?.User.introduce)
                    message.error(res.errors + '')
                }
                if (res.data) {
                    if (!res.data.UpdateUserInfo) {
                        setIntroduce(data?.User.introduce)
                    }
                }
            })
            .catch(reason => {
                setIntroduce(data?.User.introduce)
                message.error(reason + '')
            })
        setEdit(false)
    }

    return (
        <Card bordered={false} className='introduce-form' title={<span style={{float: "left"}}>个人介绍</span>}
              extra={data?.User.id === props.currentUser?.CurrentUser?.id &&
              // eslint-disable-next-line no-script-url,jsx-a11y/anchor-is-valid
              [<a href={'javascript:void(0)'} onClick={() => setEdit(true)}><IconFont type='icon-xie'/>编辑</a>]}>
            {!edit && introduce}
            {edit && <Form onFinish={onFinish}>
                <Form.Item>
                    <Input.TextArea value={introduce} className='text-a'
                                    autoSize={{minRows: 5, maxRows: 5}}
                                    onChange={e => setIntroduce(e.target.value)}
                    />
                </Form.Item>
                <Form.Item style={{float: "left"}}>
                    <Button htmlType='submit' className='btn-hollow'>保存</Button>
                    {/* eslint-disable-next-line no-script-url,jsx-a11y/anchor-is-valid */}
                    <a href={'javascript:void(0)'} className='btn-cancel'
                       onClick={() => {
                           setEdit(false)
                           setIntroduce(data?.User.introduce)
                       }}>
                        取消
                    </a>
                </Form.Item>
            </Form>}
        </Card>
    )
}