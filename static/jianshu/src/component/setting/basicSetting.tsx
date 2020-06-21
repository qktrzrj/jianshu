import React, {useState} from "react";
import {CurrentUserQuery, useUpdateUserInfoMutation, useUploadMutation} from "../../generated/graphql";
import {Avatar, Button, Input, message, Table, Upload} from "antd";
import './basicSetting.less'
import {RcCustomRequestOptions} from "antd/lib/upload/interface";

const {Column} = Table

export default function BasicSetting(props: { currentUser: CurrentUserQuery }) {

    const data = [
        {
            key: '0',
            title: 'top',
            row: 0,
        },
        {
            key: '1',
            title: '用户名',
            row: 1,
        },
        {
            key: '2',
            title: '电子邮件',
            row: 2,
        },
    ]


    const {currentUser: {CurrentUser: currentUser}} = props

    const [updateUserInfo] = useUpdateUserInfoMutation()
    const [upload] = useUploadMutation()

    const [avatar, setAvatar] = useState(currentUser.avatar)
    const [username, setUsername] = useState(currentUser.username)
    const [email, setEmail] = useState(currentUser.email)


    const onUpload = (url: string) => {
        updateUserInfo({
            variables: {avatar: url}
        })
            .then(res => {
                if (res.errors) {
                    message.error(res.errors + '')
                }
                if (res.data && res.data.UpdateUserInfo) {
                    setAvatar(url)
                } else {
                    message.error('头像更换失败')
                }
            })
            .catch(reason => message.error(reason + ''))
    }

    const customRequest = (options: RcCustomRequestOptions) => {
        upload({variables: {file: options.file}})
            .then(res => {
                if (res.errors) {
                    message.error(res.errors + '')
                }
                if (res.data) {
                    onUpload('http://localhost:8008/image/' + res.data.Upload)
                }
            })
            .catch(reason => message.error(reason + ''))
    }

    const onSave = () => {
        updateUserInfo({
            variables: {username: username, email: email}
        })
            .then(res => {
                if (res.errors) {
                    message.error(res.errors + '')
                }
                if (res.data && res.data.UpdateUserInfo) {
                    message.success('保存成功')
                } else {
                    message.error('修改用户信息失败')
                }
            })
            .catch(reason => message.error(reason + ''))
    }


    return (
        <Table dataSource={data}
               showHeader={false}
               pagination={false}
               className='basic'
               footer={() => <Button className='save' shape={"round"} type={"primary"} onClick={onSave}>保存</Button>}
        >
            <Column dataIndex='title' key='title' render={title => {
                if (title !== 'top') {
                    return <span>{title}</span>
                }
                return <Avatar className='avatar' src={avatar}/>
            }}/>
            <Column dataIndex='row' key='op' render={row => {
                switch (row) {
                    case 0: {
                        return <Upload
                            showUploadList={false}
                            action="https://upload-z2.qiniup.com/"
                            customRequest={customRequest}
                        >
                            <Button className='btn btn-hollow'>更换头像</Button>
                        </Upload>
                    }
                    case 1: {
                        return <Input className='update-input' value={username}
                                      onChange={e => setUsername(e.target.value)}/>
                    }
                    case 2: {
                        return <Input className='update-input' value={email}
                                      onChange={e => setEmail(e.target.value)}/>
                    }
                }
            }}/>
        </Table>
    )
}