import React, {useEffect} from "react";
import {Tabs} from "antd";
import {RouteComponentProps} from "react-router-dom";
import {CurrentUserQuery} from "../../generated/graphql";
import {IconFont} from "../../component/IconFont";
import './setting.less';
import BasicSetting from "../../component/setting/basicSetting";

const {TabPane} = Tabs

export default function Setting(props: RouteComponentProps & { currentUser: CurrentUserQuery }) {
    useEffect(() => {
        document.title = '设置'
    }, [])

    return (
        // @ts-ignore
        <Tabs defaultActiveKey={props.match.params.key}
              tabPosition={"left"}
              type={"card"}
              tabBarGutter={5}
              onChange={key => props.history.replace('/setting/' + key)}
        >
            <TabPane key='basic' tab={<span className='s-tab'><IconFont type='icon-shezhiwendang'/>基本设置</span>}>
                <BasicSetting {...props}/>
            </TabPane>
            {/*<TabPane key='profile' tab={<span className='s-tab'><IconFont type='icon-gerenziliao1'/>个人资料</span>}>*/}
            {/*    个人资料*/}
            {/*</TabPane>*/}
            {/*<TabPane key='misc' tab={<span className='s-tab'><IconFont type='icon-guanlipingtai'/>账号管理</span>}>*/}
            {/*    账号管理*/}
            {/*</TabPane>*/}
        </Tabs>
    )
}