import React, {useState} from "react";
import {Tabs} from "antd";
import {RouteComponentProps} from "react-router-dom";
import {CurrentUserQuery} from "../../generated/graphql";
import {IconFont} from "../../component/IconFont";
import './setting.less';
import BasicSetting from "../../component/setting/basicSetting";

const {TabPane} = Tabs

export default function Setting(props: RouteComponentProps & { currentUser: CurrentUserQuery }) {

    document.title = '设置'

    //@ts-ignore
    const [key] = useState(props.match.params.key)

    return (
        <Tabs activeKey={key}
              tabPosition={"left"}
              type={"card"}
              tabBarGutter={5}
              onChange={key => props.history.replace('/setting/' + key)}
        >
            <TabPane key='basic' tab={<span className='s-tab'><IconFont type='icon-shezhiwendang'/>基本设置</span>}>
                <BasicSetting {...props}/>
            </TabPane>
        </Tabs>
    )
}