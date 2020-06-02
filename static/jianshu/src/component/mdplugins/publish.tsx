import React from "react";
import {PluginComponent} from 'react-markdown-editor-lite';
import {IconFont} from "../IconFont";

interface PublishProps {
    State: string,
    OnClick: () => void
}


class Publish extends PluginComponent<PublishProps> {
    static pluginName = 'publish'
    static align = 'right'

    constructor(props: any) {
        super(props);

        this.handleClick = this.handleClick.bind(this);
    }

    handleClick() {
        return this.props.config.OnClick()
    }

    render() {
        return (
            <span
                className="button"
                title="发表"
                onClick={this.handleClick}
            >
             <IconFont type='icon-qianjin'/>{this.props.config.State}
            </span>
        )
    }
}

export default Publish