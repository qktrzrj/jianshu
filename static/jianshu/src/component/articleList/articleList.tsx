import React from "react";
import {List, Skeleton} from "antd";
import {Link} from "react-router-dom";
import InfiniteScroll from "react-infinite-scroller";
import {MessageOutlined, LikeOutlined, EyeOutlined} from '@ant-design/icons';
import {ListLocale} from "antd/es/list";
import './articleList.less'
import {ArticleEdge} from "../../generated/graphql";
import MarkdownIt from "markdown-it";

const IconText = ({icon, text}: any) => (
    <span>
    {React.createElement(icon, {style: {marginRight: 8}})}
        {text}
  </span>
);

export default function ArticleList(props: { fetchData: () => void, loading: boolean, hasMore: boolean | undefined, data: any[], locate?: ListLocale }) {
    const {fetchData, loading, hasMore, data, locate} = props
    const mdParser = new MarkdownIt({
        html: false,
        linkify: true,
        typographer: false,
        highlight(str, lang) {
            return '';
        },

    })
    return (
        <InfiniteScroll
            initialLoad={false}
            pageStart={0}
            loadMore={fetchData}
            hasMore={!loading && hasMore}>
            <List
                itemLayout="vertical"
                size="large"
                loading={loading}
                locale={locate}
                dataSource={data}
                renderItem={(item: ArticleEdge) => (
                    <List.Item
                        key={item.node.title}
                        actions={[
                            <IconText icon={EyeOutlined} text={item.node.ViewNum} key="list-vertical-star-o"/>,
                            <IconText icon={LikeOutlined} text={item.node.LikeNum} key="list-vertical-like-o"/>,
                            <IconText icon={MessageOutlined} text={item.node.CmtNum} key="list-vertical-message"/>,
                        ]}
                        extra={
                            item.node.cover !== '' && <img
                                width={200}
                                alt="logo"
                                src={item.node.cover}
                            />
                        }
                    >
                        <Skeleton loading={loading} active>
                            <List.Item.Meta
                                title={<Link style={{display: "flex"}}
                                             to={'/article/' + item.node.id}>{item.node.title}</Link>}
                                description={
                                    <div className='desc'
                                         dangerouslySetInnerHTML={{__html: mdParser.render(item.node.subTitle)}}/>
                                }
                            />
                        </Skeleton>
                    </List.Item>
                )}
            />
        </InfiniteScroll>
    )
}