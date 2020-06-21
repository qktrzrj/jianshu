import gql from 'graphql-tag';
import * as React from 'react';
import * as ApolloReactCommon from '@apollo/react-common';
import * as ApolloReactComponents from '@apollo/react-components';
import * as ApolloReactHoc from '@apollo/react-hoc';
import * as ApolloReactHooks from '@apollo/react-hooks';
export type Maybe<T> = T | null;
export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  /** time type */
  Time: any;
  /** Alias For String */
  NullString: any;
  /** int64 is the set of all signed 64-bit integers. Range: -9223372036854775808 through 9223372036854775807. */
  Int64: any;
  Upload: any;
};

export type Article = {
   __typename?: 'Article';
  CmtNum: Scalars['Int'];
  CommentList?: Maybe<Array<Comment>>;
  LikeNum: Scalars['Int'];
  User: User;
  ViewNum: Scalars['Int'];
  content: Scalars['String'];
  cover?: Maybe<Scalars['NullString']>;
  createdAt: Scalars['Time'];
  id: Scalars['Int'];
  state: ArticleState;
  subTitle: Scalars['String'];
  title: Scalars['String'];
  uid: Scalars['Int'];
  updatedAt: Scalars['Time'];
};

export type ArticleConnection = {
   __typename?: 'ArticleConnection';
  edges?: Maybe<Array<ArticleEdge>>;
  pageInfo: PageInfo;
  totalCount: Scalars['Int'];
};

export type ArticleEdge = {
   __typename?: 'ArticleEdge';
  cursor: Scalars['String'];
  node: Article;
};

export enum ArticleState {
  /** 已删除 */
  Deleted = 'Deleted',
  /** 草稿 */
  Draft = 'Draft',
  /** 已下线 */
  Offline = 'Offline',
  /** 已发布 */
  Online = 'Online',
  /** 未审核 */
  Unaudited = 'Unaudited',
  /** 更新未重新发布 */
  Updated = 'Updated'
}

export type Comment = {
   __typename?: 'Comment';
  User: User;
  content: Scalars['String'];
  floor: Scalars['Int'];
  id: Scalars['Int'];
  likeNum: Scalars['Int'];
  updatedAt: Scalars['Time'];
};

export enum Gender {
  Man = 'Man',
  Unknown = 'Unknown',
  Woman = 'Woman'
}


export type Msg = {
   __typename?: 'Msg';
  User: User;
  content: Scalars['String'];
  id: Scalars['Int'];
  type: MsgType;
  updatedAt: Scalars['Time'];
};

export type MsgNum = {
   __typename?: 'MsgNum';
  comment: Scalars['Int'];
  follow: Scalars['Int'];
  like: Scalars['Int'];
  reply: Scalars['Int'];
};

export enum MsgType {
  CommentMsg = 'CommentMsg',
  FollowMsg = 'FollowMsg',
  LikeMsg = 'LikeMsg',
  ReplyMsg = 'ReplyMsg'
}

export type Mutation = {
   __typename?: 'Mutation';
  AddComment: Comment;
  AddMsg: Scalars['Boolean'];
  AddReply: Reply;
  DeleteArticle: Scalars['Boolean'];
  DraftArticle: Article;
  Follow: Scalars['Boolean'];
  Like: Scalars['Boolean'];
  Logout: Scalars['Boolean'];
  NewArticle: Article;
  SignIn: User;
  SignUp: User;
  UnFollow: Scalars['Boolean'];
  Unlike: Scalars['Boolean'];
  UpdateArticle: Article;
  UpdateUserInfo: Scalars['Boolean'];
  Upload: Scalars['String'];
  ViewAdd: Scalars['Boolean'];
};


export type MutationAddCommentArgs = {
  content: Scalars['String'];
  id: Scalars['Int'];
};


export type MutationAddMsgArgs = {
  content: Scalars['String'];
  fromId: Scalars['Int'];
  toId: Scalars['Int'];
  typ: MsgType;
};


export type MutationAddReplyArgs = {
  content: Scalars['String'];
  id: Scalars['Int'];
};


export type MutationDeleteArticleArgs = {
  id: Scalars['Int'];
};


export type MutationDraftArticleArgs = {
  title: Scalars['String'];
};


export type MutationFollowArgs = {
  id: Scalars['Int'];
};


export type MutationLikeArgs = {
  id: Scalars['Int'];
  objType: ObjType;
};


export type MutationNewArticleArgs = {
  id: Scalars['Int'];
};


export type MutationSignInArgs = {
  password: Scalars['String'];
  rememberme: Scalars['Boolean'];
  username: Scalars['String'];
};


export type MutationSignUpArgs = {
  email: Scalars['String'];
  password: Scalars['String'];
  username: Scalars['String'];
};


export type MutationUnFollowArgs = {
  id: Scalars['Int'];
};


export type MutationUnlikeArgs = {
  id: Scalars['Int'];
  objType: ObjType;
};


export type MutationUpdateArticleArgs = {
  content?: Maybe<Scalars['String']>;
  cover?: Maybe<Scalars['String']>;
  id: Scalars['Int'];
  subTitle?: Maybe<Scalars['String']>;
  title?: Maybe<Scalars['String']>;
};


export type MutationUpdateUserInfoArgs = {
  avatar?: Maybe<Scalars['String']>;
  email?: Maybe<Scalars['String']>;
  gender?: Maybe<Gender>;
  introduce?: Maybe<Scalars['String']>;
  password?: Maybe<Scalars['String']>;
  username?: Maybe<Scalars['String']>;
};


export type MutationUploadArgs = {
  file: Scalars['Upload'];
};


export type MutationViewAddArgs = {
  id: Scalars['Int'];
};


export enum ObjType {
  ArticleObj = 'ArticleObj',
  CommentObj = 'CommentObj',
  ReplyObj = 'ReplyObj'
}

export type PageInfo = {
   __typename?: 'PageInfo';
  endCursor?: Maybe<Scalars['String']>;
  hasNextPage: Scalars['Boolean'];
  hasPrevPage: Scalars['Boolean'];
  pages?: Maybe<Array<Scalars['String']>>;
  startCursor?: Maybe<Scalars['String']>;
};

export type Query = {
   __typename?: 'Query';
  Article: Article;
  Articles: ArticleConnection;
  CurArticles: ArticleConnection;
  CurLikeArticles: ArticleConnection;
  CurrentUser: User;
  DownLoad: Scalars['Boolean'];
  Fans?: Maybe<Array<User>>;
  Followed?: Maybe<Array<User>>;
  HasLike: Scalars['Boolean'];
  HotArticles: ArticleConnection;
  IsFollow: Scalars['Boolean'];
  ListMsg?: Maybe<Array<Msg>>;
  MsgNum: MsgNum;
  ReplyList?: Maybe<Array<Reply>>;
  User: User;
  Users: UserConnection;
  ValidEmail: Scalars['Boolean'];
  ValidUsername: Scalars['Boolean'];
};


export type QueryArticleArgs = {
  id: Scalars['Int'];
};


export type QueryArticlesArgs = {
  after?: Maybe<Scalars['String']>;
  before?: Maybe<Scalars['String']>;
  condition?: Maybe<Scalars['String']>;
  first?: Maybe<Scalars['Int64']>;
  last?: Maybe<Scalars['Int64']>;
  uid?: Maybe<Scalars['Int']>;
};


export type QueryCurArticlesArgs = {
  after?: Maybe<Scalars['String']>;
  before?: Maybe<Scalars['String']>;
  first?: Maybe<Scalars['Int64']>;
  last?: Maybe<Scalars['Int64']>;
};


export type QueryCurLikeArticlesArgs = {
  after?: Maybe<Scalars['String']>;
  before?: Maybe<Scalars['String']>;
  first?: Maybe<Scalars['Int64']>;
  last?: Maybe<Scalars['Int64']>;
};


export type QueryDownLoadArgs = {
  fileName: Scalars['String'];
};


export type QueryFansArgs = {
  id: Scalars['Int'];
};


export type QueryFollowedArgs = {
  id: Scalars['Int'];
};


export type QueryHasLikeArgs = {
  id: Scalars['Int'];
  objType: ObjType;
};


export type QueryHotArticlesArgs = {
  after?: Maybe<Scalars['String']>;
  before?: Maybe<Scalars['String']>;
  first?: Maybe<Scalars['Int64']>;
  last?: Maybe<Scalars['Int64']>;
};


export type QueryIsFollowArgs = {
  id: Scalars['Int'];
};


export type QueryListMsgArgs = {
  typ: MsgType;
};


export type QueryReplyListArgs = {
  id: Scalars['Int'];
};


export type QueryUserArgs = {
  id: Scalars['Int'];
};


export type QueryUsersArgs = {
  after?: Maybe<Scalars['String']>;
  before?: Maybe<Scalars['String']>;
  first?: Maybe<Scalars['Int64']>;
  last?: Maybe<Scalars['Int64']>;
  username?: Maybe<Scalars['String']>;
};


export type QueryValidEmailArgs = {
  email: Scalars['String'];
};


export type QueryValidUsernameArgs = {
  username: Scalars['String'];
};

export type Reply = {
   __typename?: 'Reply';
  User: User;
  content: Scalars['String'];
  id: Scalars['Int'];
  updatedAt: Scalars['Time'];
};



export type User = {
   __typename?: 'User';
  ArticleNum: Scalars['Int'];
  FansNum: Scalars['Int'];
  FollowNum: Scalars['Int'];
  LikeNum: Scalars['Int'];
  Words: Scalars['Int'];
  avatar: Scalars['String'];
  createdAt: Scalars['Time'];
  email: Scalars['String'];
  gender: Gender;
  id: Scalars['Int'];
  introduce: Scalars['NullString'];
  root: Scalars['Boolean'];
  state: UserState;
  updatedAt: Scalars['Time'];
  username: Scalars['String'];
};

export type UserConnection = {
   __typename?: 'UserConnection';
  edges?: Maybe<Array<UserEdge>>;
  pageInfo: PageInfo;
  totalCount: Scalars['Int'];
};

export type UserEdge = {
   __typename?: 'UserEdge';
  cursor: Scalars['String'];
  node: User;
};

export enum UserState {
  Forbidden = 'Forbidden',
  Freeze = 'Freeze',
  Unsigned = 'Unsigned'
}

export type ValidUsernameQueryVariables = {
  username: Scalars['String'];
};


export type ValidUsernameQuery = (
  { __typename?: 'Query' }
  & Pick<Query, 'ValidUsername'>
);

export type ValidEmailQueryVariables = {
  email: Scalars['String'];
};


export type ValidEmailQuery = (
  { __typename?: 'Query' }
  & Pick<Query, 'ValidEmail'>
);

export type SignUpMutationVariables = {
  email: Scalars['String'];
  password: Scalars['String'];
  username: Scalars['String'];
};


export type SignUpMutation = (
  { __typename?: 'Mutation' }
  & { SignUp: (
    { __typename?: 'User' }
    & Pick<User, 'id'>
  ) }
);

export type SignInMutationVariables = {
  username: Scalars['String'];
  password: Scalars['String'];
  rememberme: Scalars['Boolean'];
};


export type SignInMutation = (
  { __typename?: 'Mutation' }
  & { SignIn: (
    { __typename?: 'User' }
    & Pick<User, 'id'>
  ) }
);

export type LogoutMutationVariables = {};


export type LogoutMutation = (
  { __typename?: 'Mutation' }
  & Pick<Mutation, 'Logout'>
);

export type UserInfoFragment = (
  { __typename?: 'User' }
  & Pick<User, 'id' | 'avatar' | 'email' | 'username' | 'introduce' | 'gender' | 'state' | 'FansNum' | 'FollowNum' | 'LikeNum' | 'Words'>
);

export type CurrentUserQueryVariables = {};


export type CurrentUserQuery = (
  { __typename?: 'Query' }
  & { CurrentUser: (
    { __typename?: 'User' }
    & UserInfoFragment
  ) }
);

export type ArticlesInfoFragment = (
  { __typename?: 'Article' }
  & Pick<Article, 'id' | 'title' | 'subTitle' | 'cover' | 'content' | 'updatedAt' | 'ViewNum' | 'LikeNum' | 'CmtNum'>
  & { User: (
    { __typename?: 'User' }
    & Pick<User, 'id' | 'username'>
  ) }
);

export type HotArticlesQueryVariables = {
  cursor?: Maybe<Scalars['String']>;
};


export type HotArticlesQuery = (
  { __typename?: 'Query' }
  & { HotArticles: (
    { __typename?: 'ArticleConnection' }
    & { edges?: Maybe<Array<(
      { __typename?: 'ArticleEdge' }
      & { node: (
        { __typename?: 'Article' }
        & ArticlesInfoFragment
      ) }
    )>>, pageInfo: (
      { __typename?: 'PageInfo' }
      & Pick<PageInfo, 'endCursor' | 'hasNextPage'>
    ) }
  ) }
);

export type DraftArticleMutationVariables = {
  title: Scalars['String'];
};


export type DraftArticleMutation = (
  { __typename?: 'Mutation' }
  & { DraftArticle: (
    { __typename?: 'Article' }
    & Pick<Article, 'id' | 'title' | 'state'>
  ) }
);

export type MyArticlesQueryVariables = {};


export type MyArticlesQuery = (
  { __typename?: 'Query' }
  & { CurArticles: (
    { __typename?: 'ArticleConnection' }
    & { edges?: Maybe<Array<(
      { __typename?: 'ArticleEdge' }
      & { node: (
        { __typename?: 'Article' }
        & Pick<Article, 'id' | 'title' | 'state'>
      ) }
    )>> }
  ) }
);

export type ArticleQueryVariables = {
  id: Scalars['Int'];
};


export type ArticleQuery = (
  { __typename?: 'Query' }
  & { Article: (
    { __typename?: 'Article' }
    & Pick<Article, 'id' | 'title' | 'subTitle' | 'content' | 'state' | 'updatedAt' | 'ViewNum' | 'LikeNum' | 'CmtNum'>
    & { User: (
      { __typename?: 'User' }
      & UserInfoFragment
    ), CommentList?: Maybe<Array<(
      { __typename?: 'Comment' }
      & Pick<Comment, 'id' | 'floor' | 'content' | 'updatedAt' | 'likeNum'>
      & { User: (
        { __typename?: 'User' }
        & Pick<User, 'id' | 'username' | 'avatar'>
      ) }
    )>> }
  ) }
);

export type DeleteArticleMutationVariables = {
  id: Scalars['Int'];
};


export type DeleteArticleMutation = (
  { __typename?: 'Mutation' }
  & Pick<Mutation, 'DeleteArticle'>
);

export type UpdateArticleMutationVariables = {
  id: Scalars['Int'];
  content?: Maybe<Scalars['String']>;
  cover?: Maybe<Scalars['String']>;
  subTitle?: Maybe<Scalars['String']>;
  title?: Maybe<Scalars['String']>;
};


export type UpdateArticleMutation = (
  { __typename?: 'Mutation' }
  & { UpdateArticle: (
    { __typename?: 'Article' }
    & Pick<Article, 'id' | 'title' | 'content' | 'state'>
  ) }
);

export type NewArticleMutationVariables = {
  id: Scalars['Int'];
};


export type NewArticleMutation = (
  { __typename?: 'Mutation' }
  & { NewArticle: (
    { __typename?: 'Article' }
    & Pick<Article, 'id' | 'state'>
  ) }
);

export type UploadMutationVariables = {
  file: Scalars['Upload'];
};


export type UploadMutation = (
  { __typename?: 'Mutation' }
  & Pick<Mutation, 'Upload'>
);

export type UserQueryVariables = {
  id: Scalars['Int'];
};


export type UserQuery = (
  { __typename?: 'Query' }
  & { User: (
    { __typename?: 'User' }
    & UserInfoFragment
  ) }
);

export type UpdateUserInfoMutationVariables = {
  username?: Maybe<Scalars['String']>;
  avatar?: Maybe<Scalars['String']>;
  email?: Maybe<Scalars['String']>;
  gender?: Maybe<Gender>;
  introduce?: Maybe<Scalars['String']>;
  password?: Maybe<Scalars['String']>;
};


export type UpdateUserInfoMutation = (
  { __typename?: 'Mutation' }
  & Pick<Mutation, 'UpdateUserInfo'>
);

export type ArticlesQueryVariables = {
  cursor?: Maybe<Scalars['String']>;
  uid?: Maybe<Scalars['Int']>;
  condition?: Maybe<Scalars['String']>;
};


export type ArticlesQuery = (
  { __typename?: 'Query' }
  & { Articles: (
    { __typename?: 'ArticleConnection' }
    & { edges?: Maybe<Array<(
      { __typename?: 'ArticleEdge' }
      & { node: (
        { __typename?: 'Article' }
        & ArticlesInfoFragment
      ) }
    )>>, pageInfo: (
      { __typename?: 'PageInfo' }
      & Pick<PageInfo, 'endCursor' | 'hasNextPage'>
    ) }
  ) }
);

export type LikeArticlesQueryVariables = {
  cursor?: Maybe<Scalars['String']>;
};


export type LikeArticlesQuery = (
  { __typename?: 'Query' }
  & { CurLikeArticles: (
    { __typename?: 'ArticleConnection' }
    & { edges?: Maybe<Array<(
      { __typename?: 'ArticleEdge' }
      & { node: (
        { __typename?: 'Article' }
        & ArticlesInfoFragment
      ) }
    )>>, pageInfo: (
      { __typename?: 'PageInfo' }
      & Pick<PageInfo, 'endCursor' | 'hasNextPage'>
    ) }
  ) }
);

export type FollowListQueryVariables = {
  id: Scalars['Int'];
};


export type FollowListQuery = (
  { __typename?: 'Query' }
  & { Followed?: Maybe<Array<(
    { __typename?: 'User' }
    & UserInfoFragment
  )>> }
);

export type IsFollowQueryVariables = {
  id: Scalars['Int'];
};


export type IsFollowQuery = (
  { __typename?: 'Query' }
  & Pick<Query, 'IsFollow'>
);

export type FollowMutationVariables = {
  id: Scalars['Int'];
};


export type FollowMutation = (
  { __typename?: 'Mutation' }
  & Pick<Mutation, 'Follow'>
);

export type UnFollowMutationVariables = {
  id: Scalars['Int'];
};


export type UnFollowMutation = (
  { __typename?: 'Mutation' }
  & Pick<Mutation, 'UnFollow'>
);

export type UsersQueryVariables = {
  username?: Maybe<Scalars['String']>;
};


export type UsersQuery = (
  { __typename?: 'Query' }
  & { Users: (
    { __typename?: 'UserConnection' }
    & { edges?: Maybe<Array<(
      { __typename?: 'UserEdge' }
      & { node: (
        { __typename?: 'User' }
        & Pick<User, 'id' | 'username' | 'avatar' | 'FansNum' | 'FollowNum' | 'LikeNum' | 'Words'>
      ) }
    )>>, pageInfo: (
      { __typename?: 'PageInfo' }
      & Pick<PageInfo, 'endCursor' | 'hasNextPage'>
    ) }
  ) }
);

export type ViewAddMutationVariables = {
  id: Scalars['Int'];
};


export type ViewAddMutation = (
  { __typename?: 'Mutation' }
  & Pick<Mutation, 'ViewAdd'>
);

export type LikeMutationVariables = {
  id: Scalars['Int'];
  objtyp: ObjType;
};


export type LikeMutation = (
  { __typename?: 'Mutation' }
  & Pick<Mutation, 'Like'>
);

export type UnLikeMutationVariables = {
  id: Scalars['Int'];
  objtyp: ObjType;
};


export type UnLikeMutation = (
  { __typename?: 'Mutation' }
  & Pick<Mutation, 'Unlike'>
);

export type HasLikeQueryVariables = {
  id: Scalars['Int'];
  objtyp: ObjType;
};


export type HasLikeQuery = (
  { __typename?: 'Query' }
  & Pick<Query, 'HasLike'>
);

export type ReplyListQueryVariables = {
  id: Scalars['Int'];
};


export type ReplyListQuery = (
  { __typename?: 'Query' }
  & { ReplyList?: Maybe<Array<(
    { __typename?: 'Reply' }
    & Pick<Reply, 'id' | 'content' | 'updatedAt'>
    & { User: (
      { __typename?: 'User' }
      & Pick<User, 'id' | 'username' | 'avatar'>
    ) }
  )>> }
);

export type AddCommentMutationVariables = {
  id: Scalars['Int'];
  content: Scalars['String'];
};


export type AddCommentMutation = (
  { __typename?: 'Mutation' }
  & { AddComment: (
    { __typename?: 'Comment' }
    & Pick<Comment, 'content' | 'floor' | 'id' | 'likeNum' | 'updatedAt'>
    & { User: (
      { __typename?: 'User' }
      & Pick<User, 'id' | 'username' | 'avatar'>
    ) }
  ) }
);

export type AddReplyMutationVariables = {
  id: Scalars['Int'];
  content: Scalars['String'];
};


export type AddReplyMutation = (
  { __typename?: 'Mutation' }
  & { AddReply: (
    { __typename?: 'Reply' }
    & Pick<Reply, 'content' | 'id' | 'updatedAt'>
    & { User: (
      { __typename?: 'User' }
      & Pick<User, 'id' | 'username' | 'avatar'>
    ) }
  ) }
);

export type AddMsgMutationVariables = {
  typ: MsgType;
  fromId: Scalars['Int'];
  toId: Scalars['Int'];
  content: Scalars['String'];
};


export type AddMsgMutation = (
  { __typename?: 'Mutation' }
  & Pick<Mutation, 'AddMsg'>
);

export type MsgNumQueryVariables = {};


export type MsgNumQuery = (
  { __typename?: 'Query' }
  & { MsgNum: (
    { __typename?: 'MsgNum' }
    & Pick<MsgNum, 'comment' | 'follow' | 'like' | 'reply'>
  ) }
);

export type ListMsgQueryVariables = {
  typ: MsgType;
};


export type ListMsgQuery = (
  { __typename?: 'Query' }
  & { ListMsg?: Maybe<Array<(
    { __typename?: 'Msg' }
    & Pick<Msg, 'content' | 'updatedAt'>
    & { User: (
      { __typename?: 'User' }
      & Pick<User, 'id' | 'username' | 'avatar'>
    ) }
  )>> }
);

export const UserInfoFragmentDoc = gql`
    fragment userInfo on User {
  id
  avatar
  email
  username
  introduce
  gender
  state
  FansNum
  FollowNum
  LikeNum
  Words
}
    `;
export const ArticlesInfoFragmentDoc = gql`
    fragment articlesInfo on Article {
  id
  title
  subTitle
  cover
  content
  updatedAt
  User {
    id
    username
  }
  ViewNum
  LikeNum
  CmtNum
}
    `;
export const ValidUsernameDocument = gql`
    query ValidUsername($username: String!) {
  ValidUsername(username: $username)
}
    `;
export type ValidUsernameComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<ValidUsernameQuery, ValidUsernameQueryVariables>, 'query'> & ({ variables: ValidUsernameQueryVariables; skip?: boolean; } | { skip: boolean; });

    export const ValidUsernameComponent = (props: ValidUsernameComponentProps) => (
      <ApolloReactComponents.Query<ValidUsernameQuery, ValidUsernameQueryVariables> query={ValidUsernameDocument} {...props} />
    );
    
export type ValidUsernameProps<TChildProps = {}, TDataName extends string = 'data'> = {
      [key in TDataName]: ApolloReactHoc.DataValue<ValidUsernameQuery, ValidUsernameQueryVariables>
    } & TChildProps;
export function withValidUsername<TProps, TChildProps = {}, TDataName extends string = 'data'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  ValidUsernameQuery,
  ValidUsernameQueryVariables,
  ValidUsernameProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withQuery<TProps, ValidUsernameQuery, ValidUsernameQueryVariables, ValidUsernameProps<TChildProps, TDataName>>(ValidUsernameDocument, {
      alias: 'validUsername',
      ...operationOptions
    });
};

/**
 * __useValidUsernameQuery__
 *
 * To run a query within a React component, call `useValidUsernameQuery` and pass it any options that fit your needs.
 * When your component renders, `useValidUsernameQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useValidUsernameQuery({
 *   variables: {
 *      username: // value for 'username'
 *   },
 * });
 */
export function useValidUsernameQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<ValidUsernameQuery, ValidUsernameQueryVariables>) {
        return ApolloReactHooks.useQuery<ValidUsernameQuery, ValidUsernameQueryVariables>(ValidUsernameDocument, baseOptions);
      }
export function useValidUsernameLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<ValidUsernameQuery, ValidUsernameQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<ValidUsernameQuery, ValidUsernameQueryVariables>(ValidUsernameDocument, baseOptions);
        }
export type ValidUsernameQueryHookResult = ReturnType<typeof useValidUsernameQuery>;
export type ValidUsernameLazyQueryHookResult = ReturnType<typeof useValidUsernameLazyQuery>;
export type ValidUsernameQueryResult = ApolloReactCommon.QueryResult<ValidUsernameQuery, ValidUsernameQueryVariables>;
export const ValidEmailDocument = gql`
    query ValidEmail($email: String!) {
  ValidEmail(email: $email)
}
    `;
export type ValidEmailComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<ValidEmailQuery, ValidEmailQueryVariables>, 'query'> & ({ variables: ValidEmailQueryVariables; skip?: boolean; } | { skip: boolean; });

    export const ValidEmailComponent = (props: ValidEmailComponentProps) => (
      <ApolloReactComponents.Query<ValidEmailQuery, ValidEmailQueryVariables> query={ValidEmailDocument} {...props} />
    );
    
export type ValidEmailProps<TChildProps = {}, TDataName extends string = 'data'> = {
      [key in TDataName]: ApolloReactHoc.DataValue<ValidEmailQuery, ValidEmailQueryVariables>
    } & TChildProps;
export function withValidEmail<TProps, TChildProps = {}, TDataName extends string = 'data'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  ValidEmailQuery,
  ValidEmailQueryVariables,
  ValidEmailProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withQuery<TProps, ValidEmailQuery, ValidEmailQueryVariables, ValidEmailProps<TChildProps, TDataName>>(ValidEmailDocument, {
      alias: 'validEmail',
      ...operationOptions
    });
};

/**
 * __useValidEmailQuery__
 *
 * To run a query within a React component, call `useValidEmailQuery` and pass it any options that fit your needs.
 * When your component renders, `useValidEmailQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useValidEmailQuery({
 *   variables: {
 *      email: // value for 'email'
 *   },
 * });
 */
export function useValidEmailQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<ValidEmailQuery, ValidEmailQueryVariables>) {
        return ApolloReactHooks.useQuery<ValidEmailQuery, ValidEmailQueryVariables>(ValidEmailDocument, baseOptions);
      }
export function useValidEmailLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<ValidEmailQuery, ValidEmailQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<ValidEmailQuery, ValidEmailQueryVariables>(ValidEmailDocument, baseOptions);
        }
export type ValidEmailQueryHookResult = ReturnType<typeof useValidEmailQuery>;
export type ValidEmailLazyQueryHookResult = ReturnType<typeof useValidEmailLazyQuery>;
export type ValidEmailQueryResult = ApolloReactCommon.QueryResult<ValidEmailQuery, ValidEmailQueryVariables>;
export const SignUpDocument = gql`
    mutation SignUp($email: String!, $password: String!, $username: String!) {
  SignUp(email: $email, password: $password, username: $username) {
    id
  }
}
    `;
export type SignUpMutationFn = ApolloReactCommon.MutationFunction<SignUpMutation, SignUpMutationVariables>;
export type SignUpComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<SignUpMutation, SignUpMutationVariables>, 'mutation'>;

    export const SignUpComponent = (props: SignUpComponentProps) => (
      <ApolloReactComponents.Mutation<SignUpMutation, SignUpMutationVariables> mutation={SignUpDocument} {...props} />
    );
    
export type SignUpProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<SignUpMutation, SignUpMutationVariables>
    } & TChildProps;
export function withSignUp<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  SignUpMutation,
  SignUpMutationVariables,
  SignUpProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, SignUpMutation, SignUpMutationVariables, SignUpProps<TChildProps, TDataName>>(SignUpDocument, {
      alias: 'signUp',
      ...operationOptions
    });
};

/**
 * __useSignUpMutation__
 *
 * To run a mutation, you first call `useSignUpMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useSignUpMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [signUpMutation, { data, loading, error }] = useSignUpMutation({
 *   variables: {
 *      email: // value for 'email'
 *      password: // value for 'password'
 *      username: // value for 'username'
 *   },
 * });
 */
export function useSignUpMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<SignUpMutation, SignUpMutationVariables>) {
        return ApolloReactHooks.useMutation<SignUpMutation, SignUpMutationVariables>(SignUpDocument, baseOptions);
      }
export type SignUpMutationHookResult = ReturnType<typeof useSignUpMutation>;
export type SignUpMutationResult = ApolloReactCommon.MutationResult<SignUpMutation>;
export type SignUpMutationOptions = ApolloReactCommon.BaseMutationOptions<SignUpMutation, SignUpMutationVariables>;
export const SignInDocument = gql`
    mutation SignIn($username: String!, $password: String!, $rememberme: Boolean!) {
  SignIn(username: $username, password: $password, rememberme: $rememberme) {
    id
  }
}
    `;
export type SignInMutationFn = ApolloReactCommon.MutationFunction<SignInMutation, SignInMutationVariables>;
export type SignInComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<SignInMutation, SignInMutationVariables>, 'mutation'>;

    export const SignInComponent = (props: SignInComponentProps) => (
      <ApolloReactComponents.Mutation<SignInMutation, SignInMutationVariables> mutation={SignInDocument} {...props} />
    );
    
export type SignInProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<SignInMutation, SignInMutationVariables>
    } & TChildProps;
export function withSignIn<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  SignInMutation,
  SignInMutationVariables,
  SignInProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, SignInMutation, SignInMutationVariables, SignInProps<TChildProps, TDataName>>(SignInDocument, {
      alias: 'signIn',
      ...operationOptions
    });
};

/**
 * __useSignInMutation__
 *
 * To run a mutation, you first call `useSignInMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useSignInMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [signInMutation, { data, loading, error }] = useSignInMutation({
 *   variables: {
 *      username: // value for 'username'
 *      password: // value for 'password'
 *      rememberme: // value for 'rememberme'
 *   },
 * });
 */
export function useSignInMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<SignInMutation, SignInMutationVariables>) {
        return ApolloReactHooks.useMutation<SignInMutation, SignInMutationVariables>(SignInDocument, baseOptions);
      }
export type SignInMutationHookResult = ReturnType<typeof useSignInMutation>;
export type SignInMutationResult = ApolloReactCommon.MutationResult<SignInMutation>;
export type SignInMutationOptions = ApolloReactCommon.BaseMutationOptions<SignInMutation, SignInMutationVariables>;
export const LogoutDocument = gql`
    mutation Logout {
  Logout
}
    `;
export type LogoutMutationFn = ApolloReactCommon.MutationFunction<LogoutMutation, LogoutMutationVariables>;
export type LogoutComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<LogoutMutation, LogoutMutationVariables>, 'mutation'>;

    export const LogoutComponent = (props: LogoutComponentProps) => (
      <ApolloReactComponents.Mutation<LogoutMutation, LogoutMutationVariables> mutation={LogoutDocument} {...props} />
    );
    
export type LogoutProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<LogoutMutation, LogoutMutationVariables>
    } & TChildProps;
export function withLogout<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  LogoutMutation,
  LogoutMutationVariables,
  LogoutProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, LogoutMutation, LogoutMutationVariables, LogoutProps<TChildProps, TDataName>>(LogoutDocument, {
      alias: 'logout',
      ...operationOptions
    });
};

/**
 * __useLogoutMutation__
 *
 * To run a mutation, you first call `useLogoutMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useLogoutMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [logoutMutation, { data, loading, error }] = useLogoutMutation({
 *   variables: {
 *   },
 * });
 */
export function useLogoutMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<LogoutMutation, LogoutMutationVariables>) {
        return ApolloReactHooks.useMutation<LogoutMutation, LogoutMutationVariables>(LogoutDocument, baseOptions);
      }
export type LogoutMutationHookResult = ReturnType<typeof useLogoutMutation>;
export type LogoutMutationResult = ApolloReactCommon.MutationResult<LogoutMutation>;
export type LogoutMutationOptions = ApolloReactCommon.BaseMutationOptions<LogoutMutation, LogoutMutationVariables>;
export const CurrentUserDocument = gql`
    query CurrentUser {
  CurrentUser {
    ...userInfo
  }
}
    ${UserInfoFragmentDoc}`;
export type CurrentUserComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<CurrentUserQuery, CurrentUserQueryVariables>, 'query'>;

    export const CurrentUserComponent = (props: CurrentUserComponentProps) => (
      <ApolloReactComponents.Query<CurrentUserQuery, CurrentUserQueryVariables> query={CurrentUserDocument} {...props} />
    );
    
export type CurrentUserProps<TChildProps = {}, TDataName extends string = 'data'> = {
      [key in TDataName]: ApolloReactHoc.DataValue<CurrentUserQuery, CurrentUserQueryVariables>
    } & TChildProps;
export function withCurrentUser<TProps, TChildProps = {}, TDataName extends string = 'data'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  CurrentUserQuery,
  CurrentUserQueryVariables,
  CurrentUserProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withQuery<TProps, CurrentUserQuery, CurrentUserQueryVariables, CurrentUserProps<TChildProps, TDataName>>(CurrentUserDocument, {
      alias: 'currentUser',
      ...operationOptions
    });
};

/**
 * __useCurrentUserQuery__
 *
 * To run a query within a React component, call `useCurrentUserQuery` and pass it any options that fit your needs.
 * When your component renders, `useCurrentUserQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useCurrentUserQuery({
 *   variables: {
 *   },
 * });
 */
export function useCurrentUserQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<CurrentUserQuery, CurrentUserQueryVariables>) {
        return ApolloReactHooks.useQuery<CurrentUserQuery, CurrentUserQueryVariables>(CurrentUserDocument, baseOptions);
      }
export function useCurrentUserLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<CurrentUserQuery, CurrentUserQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<CurrentUserQuery, CurrentUserQueryVariables>(CurrentUserDocument, baseOptions);
        }
export type CurrentUserQueryHookResult = ReturnType<typeof useCurrentUserQuery>;
export type CurrentUserLazyQueryHookResult = ReturnType<typeof useCurrentUserLazyQuery>;
export type CurrentUserQueryResult = ApolloReactCommon.QueryResult<CurrentUserQuery, CurrentUserQueryVariables>;
export const HotArticlesDocument = gql`
    query HotArticles($cursor: String) {
  HotArticles(first: 10, after: $cursor) {
    edges {
      node {
        ...articlesInfo
      }
    }
    pageInfo {
      endCursor
      hasNextPage
    }
  }
}
    ${ArticlesInfoFragmentDoc}`;
export type HotArticlesComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<HotArticlesQuery, HotArticlesQueryVariables>, 'query'>;

    export const HotArticlesComponent = (props: HotArticlesComponentProps) => (
      <ApolloReactComponents.Query<HotArticlesQuery, HotArticlesQueryVariables> query={HotArticlesDocument} {...props} />
    );
    
export type HotArticlesProps<TChildProps = {}, TDataName extends string = 'data'> = {
      [key in TDataName]: ApolloReactHoc.DataValue<HotArticlesQuery, HotArticlesQueryVariables>
    } & TChildProps;
export function withHotArticles<TProps, TChildProps = {}, TDataName extends string = 'data'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  HotArticlesQuery,
  HotArticlesQueryVariables,
  HotArticlesProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withQuery<TProps, HotArticlesQuery, HotArticlesQueryVariables, HotArticlesProps<TChildProps, TDataName>>(HotArticlesDocument, {
      alias: 'hotArticles',
      ...operationOptions
    });
};

/**
 * __useHotArticlesQuery__
 *
 * To run a query within a React component, call `useHotArticlesQuery` and pass it any options that fit your needs.
 * When your component renders, `useHotArticlesQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useHotArticlesQuery({
 *   variables: {
 *      cursor: // value for 'cursor'
 *   },
 * });
 */
export function useHotArticlesQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<HotArticlesQuery, HotArticlesQueryVariables>) {
        return ApolloReactHooks.useQuery<HotArticlesQuery, HotArticlesQueryVariables>(HotArticlesDocument, baseOptions);
      }
export function useHotArticlesLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<HotArticlesQuery, HotArticlesQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<HotArticlesQuery, HotArticlesQueryVariables>(HotArticlesDocument, baseOptions);
        }
export type HotArticlesQueryHookResult = ReturnType<typeof useHotArticlesQuery>;
export type HotArticlesLazyQueryHookResult = ReturnType<typeof useHotArticlesLazyQuery>;
export type HotArticlesQueryResult = ApolloReactCommon.QueryResult<HotArticlesQuery, HotArticlesQueryVariables>;
export const DraftArticleDocument = gql`
    mutation DraftArticle($title: String!) {
  DraftArticle(title: $title) {
    id
    title
    state
  }
}
    `;
export type DraftArticleMutationFn = ApolloReactCommon.MutationFunction<DraftArticleMutation, DraftArticleMutationVariables>;
export type DraftArticleComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<DraftArticleMutation, DraftArticleMutationVariables>, 'mutation'>;

    export const DraftArticleComponent = (props: DraftArticleComponentProps) => (
      <ApolloReactComponents.Mutation<DraftArticleMutation, DraftArticleMutationVariables> mutation={DraftArticleDocument} {...props} />
    );
    
export type DraftArticleProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<DraftArticleMutation, DraftArticleMutationVariables>
    } & TChildProps;
export function withDraftArticle<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  DraftArticleMutation,
  DraftArticleMutationVariables,
  DraftArticleProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, DraftArticleMutation, DraftArticleMutationVariables, DraftArticleProps<TChildProps, TDataName>>(DraftArticleDocument, {
      alias: 'draftArticle',
      ...operationOptions
    });
};

/**
 * __useDraftArticleMutation__
 *
 * To run a mutation, you first call `useDraftArticleMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDraftArticleMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [draftArticleMutation, { data, loading, error }] = useDraftArticleMutation({
 *   variables: {
 *      title: // value for 'title'
 *   },
 * });
 */
export function useDraftArticleMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<DraftArticleMutation, DraftArticleMutationVariables>) {
        return ApolloReactHooks.useMutation<DraftArticleMutation, DraftArticleMutationVariables>(DraftArticleDocument, baseOptions);
      }
export type DraftArticleMutationHookResult = ReturnType<typeof useDraftArticleMutation>;
export type DraftArticleMutationResult = ApolloReactCommon.MutationResult<DraftArticleMutation>;
export type DraftArticleMutationOptions = ApolloReactCommon.BaseMutationOptions<DraftArticleMutation, DraftArticleMutationVariables>;
export const MyArticlesDocument = gql`
    query MyArticles {
  CurArticles {
    edges {
      node {
        id
        title
        state
      }
    }
  }
}
    `;
export type MyArticlesComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<MyArticlesQuery, MyArticlesQueryVariables>, 'query'>;

    export const MyArticlesComponent = (props: MyArticlesComponentProps) => (
      <ApolloReactComponents.Query<MyArticlesQuery, MyArticlesQueryVariables> query={MyArticlesDocument} {...props} />
    );
    
export type MyArticlesProps<TChildProps = {}, TDataName extends string = 'data'> = {
      [key in TDataName]: ApolloReactHoc.DataValue<MyArticlesQuery, MyArticlesQueryVariables>
    } & TChildProps;
export function withMyArticles<TProps, TChildProps = {}, TDataName extends string = 'data'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  MyArticlesQuery,
  MyArticlesQueryVariables,
  MyArticlesProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withQuery<TProps, MyArticlesQuery, MyArticlesQueryVariables, MyArticlesProps<TChildProps, TDataName>>(MyArticlesDocument, {
      alias: 'myArticles',
      ...operationOptions
    });
};

/**
 * __useMyArticlesQuery__
 *
 * To run a query within a React component, call `useMyArticlesQuery` and pass it any options that fit your needs.
 * When your component renders, `useMyArticlesQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useMyArticlesQuery({
 *   variables: {
 *   },
 * });
 */
export function useMyArticlesQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<MyArticlesQuery, MyArticlesQueryVariables>) {
        return ApolloReactHooks.useQuery<MyArticlesQuery, MyArticlesQueryVariables>(MyArticlesDocument, baseOptions);
      }
export function useMyArticlesLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<MyArticlesQuery, MyArticlesQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<MyArticlesQuery, MyArticlesQueryVariables>(MyArticlesDocument, baseOptions);
        }
export type MyArticlesQueryHookResult = ReturnType<typeof useMyArticlesQuery>;
export type MyArticlesLazyQueryHookResult = ReturnType<typeof useMyArticlesLazyQuery>;
export type MyArticlesQueryResult = ApolloReactCommon.QueryResult<MyArticlesQuery, MyArticlesQueryVariables>;
export const ArticleDocument = gql`
    query Article($id: Int!) {
  Article(id: $id) {
    id
    title
    subTitle
    content
    state
    updatedAt
    ViewNum
    LikeNum
    CmtNum
    User {
      ...userInfo
    }
    CommentList {
      id
      floor
      content
      updatedAt
      likeNum
      User {
        id
        username
        avatar
      }
    }
  }
}
    ${UserInfoFragmentDoc}`;
export type ArticleComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<ArticleQuery, ArticleQueryVariables>, 'query'> & ({ variables: ArticleQueryVariables; skip?: boolean; } | { skip: boolean; });

    export const ArticleComponent = (props: ArticleComponentProps) => (
      <ApolloReactComponents.Query<ArticleQuery, ArticleQueryVariables> query={ArticleDocument} {...props} />
    );
    
export type ArticleProps<TChildProps = {}, TDataName extends string = 'data'> = {
      [key in TDataName]: ApolloReactHoc.DataValue<ArticleQuery, ArticleQueryVariables>
    } & TChildProps;
export function withArticle<TProps, TChildProps = {}, TDataName extends string = 'data'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  ArticleQuery,
  ArticleQueryVariables,
  ArticleProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withQuery<TProps, ArticleQuery, ArticleQueryVariables, ArticleProps<TChildProps, TDataName>>(ArticleDocument, {
      alias: 'article',
      ...operationOptions
    });
};

/**
 * __useArticleQuery__
 *
 * To run a query within a React component, call `useArticleQuery` and pass it any options that fit your needs.
 * When your component renders, `useArticleQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useArticleQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useArticleQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<ArticleQuery, ArticleQueryVariables>) {
        return ApolloReactHooks.useQuery<ArticleQuery, ArticleQueryVariables>(ArticleDocument, baseOptions);
      }
export function useArticleLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<ArticleQuery, ArticleQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<ArticleQuery, ArticleQueryVariables>(ArticleDocument, baseOptions);
        }
export type ArticleQueryHookResult = ReturnType<typeof useArticleQuery>;
export type ArticleLazyQueryHookResult = ReturnType<typeof useArticleLazyQuery>;
export type ArticleQueryResult = ApolloReactCommon.QueryResult<ArticleQuery, ArticleQueryVariables>;
export const DeleteArticleDocument = gql`
    mutation DeleteArticle($id: Int!) {
  DeleteArticle(id: $id)
}
    `;
export type DeleteArticleMutationFn = ApolloReactCommon.MutationFunction<DeleteArticleMutation, DeleteArticleMutationVariables>;
export type DeleteArticleComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<DeleteArticleMutation, DeleteArticleMutationVariables>, 'mutation'>;

    export const DeleteArticleComponent = (props: DeleteArticleComponentProps) => (
      <ApolloReactComponents.Mutation<DeleteArticleMutation, DeleteArticleMutationVariables> mutation={DeleteArticleDocument} {...props} />
    );
    
export type DeleteArticleProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<DeleteArticleMutation, DeleteArticleMutationVariables>
    } & TChildProps;
export function withDeleteArticle<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  DeleteArticleMutation,
  DeleteArticleMutationVariables,
  DeleteArticleProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, DeleteArticleMutation, DeleteArticleMutationVariables, DeleteArticleProps<TChildProps, TDataName>>(DeleteArticleDocument, {
      alias: 'deleteArticle',
      ...operationOptions
    });
};

/**
 * __useDeleteArticleMutation__
 *
 * To run a mutation, you first call `useDeleteArticleMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteArticleMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteArticleMutation, { data, loading, error }] = useDeleteArticleMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useDeleteArticleMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<DeleteArticleMutation, DeleteArticleMutationVariables>) {
        return ApolloReactHooks.useMutation<DeleteArticleMutation, DeleteArticleMutationVariables>(DeleteArticleDocument, baseOptions);
      }
export type DeleteArticleMutationHookResult = ReturnType<typeof useDeleteArticleMutation>;
export type DeleteArticleMutationResult = ApolloReactCommon.MutationResult<DeleteArticleMutation>;
export type DeleteArticleMutationOptions = ApolloReactCommon.BaseMutationOptions<DeleteArticleMutation, DeleteArticleMutationVariables>;
export const UpdateArticleDocument = gql`
    mutation UpdateArticle($id: Int!, $content: String, $cover: String, $subTitle: String, $title: String) {
  UpdateArticle(id: $id, content: $content, cover: $cover, subTitle: $subTitle, title: $title) {
    id
    title
    content
    state
  }
}
    `;
export type UpdateArticleMutationFn = ApolloReactCommon.MutationFunction<UpdateArticleMutation, UpdateArticleMutationVariables>;
export type UpdateArticleComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<UpdateArticleMutation, UpdateArticleMutationVariables>, 'mutation'>;

    export const UpdateArticleComponent = (props: UpdateArticleComponentProps) => (
      <ApolloReactComponents.Mutation<UpdateArticleMutation, UpdateArticleMutationVariables> mutation={UpdateArticleDocument} {...props} />
    );
    
export type UpdateArticleProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<UpdateArticleMutation, UpdateArticleMutationVariables>
    } & TChildProps;
export function withUpdateArticle<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  UpdateArticleMutation,
  UpdateArticleMutationVariables,
  UpdateArticleProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, UpdateArticleMutation, UpdateArticleMutationVariables, UpdateArticleProps<TChildProps, TDataName>>(UpdateArticleDocument, {
      alias: 'updateArticle',
      ...operationOptions
    });
};

/**
 * __useUpdateArticleMutation__
 *
 * To run a mutation, you first call `useUpdateArticleMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateArticleMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateArticleMutation, { data, loading, error }] = useUpdateArticleMutation({
 *   variables: {
 *      id: // value for 'id'
 *      content: // value for 'content'
 *      cover: // value for 'cover'
 *      subTitle: // value for 'subTitle'
 *      title: // value for 'title'
 *   },
 * });
 */
export function useUpdateArticleMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<UpdateArticleMutation, UpdateArticleMutationVariables>) {
        return ApolloReactHooks.useMutation<UpdateArticleMutation, UpdateArticleMutationVariables>(UpdateArticleDocument, baseOptions);
      }
export type UpdateArticleMutationHookResult = ReturnType<typeof useUpdateArticleMutation>;
export type UpdateArticleMutationResult = ApolloReactCommon.MutationResult<UpdateArticleMutation>;
export type UpdateArticleMutationOptions = ApolloReactCommon.BaseMutationOptions<UpdateArticleMutation, UpdateArticleMutationVariables>;
export const NewArticleDocument = gql`
    mutation NewArticle($id: Int!) {
  NewArticle(id: $id) {
    id
    state
  }
}
    `;
export type NewArticleMutationFn = ApolloReactCommon.MutationFunction<NewArticleMutation, NewArticleMutationVariables>;
export type NewArticleComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<NewArticleMutation, NewArticleMutationVariables>, 'mutation'>;

    export const NewArticleComponent = (props: NewArticleComponentProps) => (
      <ApolloReactComponents.Mutation<NewArticleMutation, NewArticleMutationVariables> mutation={NewArticleDocument} {...props} />
    );
    
export type NewArticleProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<NewArticleMutation, NewArticleMutationVariables>
    } & TChildProps;
export function withNewArticle<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  NewArticleMutation,
  NewArticleMutationVariables,
  NewArticleProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, NewArticleMutation, NewArticleMutationVariables, NewArticleProps<TChildProps, TDataName>>(NewArticleDocument, {
      alias: 'newArticle',
      ...operationOptions
    });
};

/**
 * __useNewArticleMutation__
 *
 * To run a mutation, you first call `useNewArticleMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useNewArticleMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [newArticleMutation, { data, loading, error }] = useNewArticleMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useNewArticleMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<NewArticleMutation, NewArticleMutationVariables>) {
        return ApolloReactHooks.useMutation<NewArticleMutation, NewArticleMutationVariables>(NewArticleDocument, baseOptions);
      }
export type NewArticleMutationHookResult = ReturnType<typeof useNewArticleMutation>;
export type NewArticleMutationResult = ApolloReactCommon.MutationResult<NewArticleMutation>;
export type NewArticleMutationOptions = ApolloReactCommon.BaseMutationOptions<NewArticleMutation, NewArticleMutationVariables>;
export const UploadDocument = gql`
    mutation Upload($file: Upload!) {
  Upload(file: $file)
}
    `;
export type UploadMutationFn = ApolloReactCommon.MutationFunction<UploadMutation, UploadMutationVariables>;
export type UploadComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<UploadMutation, UploadMutationVariables>, 'mutation'>;

    export const UploadComponent = (props: UploadComponentProps) => (
      <ApolloReactComponents.Mutation<UploadMutation, UploadMutationVariables> mutation={UploadDocument} {...props} />
    );
    
export type UploadProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<UploadMutation, UploadMutationVariables>
    } & TChildProps;
export function withUpload<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  UploadMutation,
  UploadMutationVariables,
  UploadProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, UploadMutation, UploadMutationVariables, UploadProps<TChildProps, TDataName>>(UploadDocument, {
      alias: 'upload',
      ...operationOptions
    });
};

/**
 * __useUploadMutation__
 *
 * To run a mutation, you first call `useUploadMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUploadMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [uploadMutation, { data, loading, error }] = useUploadMutation({
 *   variables: {
 *      file: // value for 'file'
 *   },
 * });
 */
export function useUploadMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<UploadMutation, UploadMutationVariables>) {
        return ApolloReactHooks.useMutation<UploadMutation, UploadMutationVariables>(UploadDocument, baseOptions);
      }
export type UploadMutationHookResult = ReturnType<typeof useUploadMutation>;
export type UploadMutationResult = ApolloReactCommon.MutationResult<UploadMutation>;
export type UploadMutationOptions = ApolloReactCommon.BaseMutationOptions<UploadMutation, UploadMutationVariables>;
export const UserDocument = gql`
    query User($id: Int!) {
  User(id: $id) {
    ...userInfo
  }
}
    ${UserInfoFragmentDoc}`;
export type UserComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<UserQuery, UserQueryVariables>, 'query'> & ({ variables: UserQueryVariables; skip?: boolean; } | { skip: boolean; });

    export const UserComponent = (props: UserComponentProps) => (
      <ApolloReactComponents.Query<UserQuery, UserQueryVariables> query={UserDocument} {...props} />
    );
    
export type UserProps<TChildProps = {}, TDataName extends string = 'data'> = {
      [key in TDataName]: ApolloReactHoc.DataValue<UserQuery, UserQueryVariables>
    } & TChildProps;
export function withUser<TProps, TChildProps = {}, TDataName extends string = 'data'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  UserQuery,
  UserQueryVariables,
  UserProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withQuery<TProps, UserQuery, UserQueryVariables, UserProps<TChildProps, TDataName>>(UserDocument, {
      alias: 'user',
      ...operationOptions
    });
};

/**
 * __useUserQuery__
 *
 * To run a query within a React component, call `useUserQuery` and pass it any options that fit your needs.
 * When your component renders, `useUserQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useUserQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useUserQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<UserQuery, UserQueryVariables>) {
        return ApolloReactHooks.useQuery<UserQuery, UserQueryVariables>(UserDocument, baseOptions);
      }
export function useUserLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<UserQuery, UserQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<UserQuery, UserQueryVariables>(UserDocument, baseOptions);
        }
export type UserQueryHookResult = ReturnType<typeof useUserQuery>;
export type UserLazyQueryHookResult = ReturnType<typeof useUserLazyQuery>;
export type UserQueryResult = ApolloReactCommon.QueryResult<UserQuery, UserQueryVariables>;
export const UpdateUserInfoDocument = gql`
    mutation UpdateUserInfo($username: String = null, $avatar: String = null, $email: String = null, $gender: Gender = null, $introduce: String = null, $password: String = null) {
  UpdateUserInfo(username: $username, avatar: $avatar, email: $email, gender: $gender, introduce: $introduce, password: $password)
}
    `;
export type UpdateUserInfoMutationFn = ApolloReactCommon.MutationFunction<UpdateUserInfoMutation, UpdateUserInfoMutationVariables>;
export type UpdateUserInfoComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<UpdateUserInfoMutation, UpdateUserInfoMutationVariables>, 'mutation'>;

    export const UpdateUserInfoComponent = (props: UpdateUserInfoComponentProps) => (
      <ApolloReactComponents.Mutation<UpdateUserInfoMutation, UpdateUserInfoMutationVariables> mutation={UpdateUserInfoDocument} {...props} />
    );
    
export type UpdateUserInfoProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<UpdateUserInfoMutation, UpdateUserInfoMutationVariables>
    } & TChildProps;
export function withUpdateUserInfo<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  UpdateUserInfoMutation,
  UpdateUserInfoMutationVariables,
  UpdateUserInfoProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, UpdateUserInfoMutation, UpdateUserInfoMutationVariables, UpdateUserInfoProps<TChildProps, TDataName>>(UpdateUserInfoDocument, {
      alias: 'updateUserInfo',
      ...operationOptions
    });
};

/**
 * __useUpdateUserInfoMutation__
 *
 * To run a mutation, you first call `useUpdateUserInfoMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateUserInfoMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateUserInfoMutation, { data, loading, error }] = useUpdateUserInfoMutation({
 *   variables: {
 *      username: // value for 'username'
 *      avatar: // value for 'avatar'
 *      email: // value for 'email'
 *      gender: // value for 'gender'
 *      introduce: // value for 'introduce'
 *      password: // value for 'password'
 *   },
 * });
 */
export function useUpdateUserInfoMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<UpdateUserInfoMutation, UpdateUserInfoMutationVariables>) {
        return ApolloReactHooks.useMutation<UpdateUserInfoMutation, UpdateUserInfoMutationVariables>(UpdateUserInfoDocument, baseOptions);
      }
export type UpdateUserInfoMutationHookResult = ReturnType<typeof useUpdateUserInfoMutation>;
export type UpdateUserInfoMutationResult = ApolloReactCommon.MutationResult<UpdateUserInfoMutation>;
export type UpdateUserInfoMutationOptions = ApolloReactCommon.BaseMutationOptions<UpdateUserInfoMutation, UpdateUserInfoMutationVariables>;
export const ArticlesDocument = gql`
    query Articles($cursor: String = null, $uid: Int = null, $condition: String = null) {
  Articles(first: 10, after: $cursor, uid: $uid, condition: $condition) {
    edges {
      node {
        ...articlesInfo
      }
    }
    pageInfo {
      endCursor
      hasNextPage
    }
  }
}
    ${ArticlesInfoFragmentDoc}`;
export type ArticlesComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<ArticlesQuery, ArticlesQueryVariables>, 'query'>;

    export const ArticlesComponent = (props: ArticlesComponentProps) => (
      <ApolloReactComponents.Query<ArticlesQuery, ArticlesQueryVariables> query={ArticlesDocument} {...props} />
    );
    
export type ArticlesProps<TChildProps = {}, TDataName extends string = 'data'> = {
      [key in TDataName]: ApolloReactHoc.DataValue<ArticlesQuery, ArticlesQueryVariables>
    } & TChildProps;
export function withArticles<TProps, TChildProps = {}, TDataName extends string = 'data'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  ArticlesQuery,
  ArticlesQueryVariables,
  ArticlesProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withQuery<TProps, ArticlesQuery, ArticlesQueryVariables, ArticlesProps<TChildProps, TDataName>>(ArticlesDocument, {
      alias: 'articles',
      ...operationOptions
    });
};

/**
 * __useArticlesQuery__
 *
 * To run a query within a React component, call `useArticlesQuery` and pass it any options that fit your needs.
 * When your component renders, `useArticlesQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useArticlesQuery({
 *   variables: {
 *      cursor: // value for 'cursor'
 *      uid: // value for 'uid'
 *      condition: // value for 'condition'
 *   },
 * });
 */
export function useArticlesQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<ArticlesQuery, ArticlesQueryVariables>) {
        return ApolloReactHooks.useQuery<ArticlesQuery, ArticlesQueryVariables>(ArticlesDocument, baseOptions);
      }
export function useArticlesLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<ArticlesQuery, ArticlesQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<ArticlesQuery, ArticlesQueryVariables>(ArticlesDocument, baseOptions);
        }
export type ArticlesQueryHookResult = ReturnType<typeof useArticlesQuery>;
export type ArticlesLazyQueryHookResult = ReturnType<typeof useArticlesLazyQuery>;
export type ArticlesQueryResult = ApolloReactCommon.QueryResult<ArticlesQuery, ArticlesQueryVariables>;
export const LikeArticlesDocument = gql`
    query LikeArticles($cursor: String = null) {
  CurLikeArticles(first: 10, after: $cursor) {
    edges {
      node {
        ...articlesInfo
      }
    }
    pageInfo {
      endCursor
      hasNextPage
    }
  }
}
    ${ArticlesInfoFragmentDoc}`;
export type LikeArticlesComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<LikeArticlesQuery, LikeArticlesQueryVariables>, 'query'>;

    export const LikeArticlesComponent = (props: LikeArticlesComponentProps) => (
      <ApolloReactComponents.Query<LikeArticlesQuery, LikeArticlesQueryVariables> query={LikeArticlesDocument} {...props} />
    );
    
export type LikeArticlesProps<TChildProps = {}, TDataName extends string = 'data'> = {
      [key in TDataName]: ApolloReactHoc.DataValue<LikeArticlesQuery, LikeArticlesQueryVariables>
    } & TChildProps;
export function withLikeArticles<TProps, TChildProps = {}, TDataName extends string = 'data'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  LikeArticlesQuery,
  LikeArticlesQueryVariables,
  LikeArticlesProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withQuery<TProps, LikeArticlesQuery, LikeArticlesQueryVariables, LikeArticlesProps<TChildProps, TDataName>>(LikeArticlesDocument, {
      alias: 'likeArticles',
      ...operationOptions
    });
};

/**
 * __useLikeArticlesQuery__
 *
 * To run a query within a React component, call `useLikeArticlesQuery` and pass it any options that fit your needs.
 * When your component renders, `useLikeArticlesQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useLikeArticlesQuery({
 *   variables: {
 *      cursor: // value for 'cursor'
 *   },
 * });
 */
export function useLikeArticlesQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<LikeArticlesQuery, LikeArticlesQueryVariables>) {
        return ApolloReactHooks.useQuery<LikeArticlesQuery, LikeArticlesQueryVariables>(LikeArticlesDocument, baseOptions);
      }
export function useLikeArticlesLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<LikeArticlesQuery, LikeArticlesQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<LikeArticlesQuery, LikeArticlesQueryVariables>(LikeArticlesDocument, baseOptions);
        }
export type LikeArticlesQueryHookResult = ReturnType<typeof useLikeArticlesQuery>;
export type LikeArticlesLazyQueryHookResult = ReturnType<typeof useLikeArticlesLazyQuery>;
export type LikeArticlesQueryResult = ApolloReactCommon.QueryResult<LikeArticlesQuery, LikeArticlesQueryVariables>;
export const FollowListDocument = gql`
    query FollowList($id: Int!) {
  Followed(id: $id) {
    ...userInfo
  }
}
    ${UserInfoFragmentDoc}`;
export type FollowListComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<FollowListQuery, FollowListQueryVariables>, 'query'> & ({ variables: FollowListQueryVariables; skip?: boolean; } | { skip: boolean; });

    export const FollowListComponent = (props: FollowListComponentProps) => (
      <ApolloReactComponents.Query<FollowListQuery, FollowListQueryVariables> query={FollowListDocument} {...props} />
    );
    
export type FollowListProps<TChildProps = {}, TDataName extends string = 'data'> = {
      [key in TDataName]: ApolloReactHoc.DataValue<FollowListQuery, FollowListQueryVariables>
    } & TChildProps;
export function withFollowList<TProps, TChildProps = {}, TDataName extends string = 'data'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  FollowListQuery,
  FollowListQueryVariables,
  FollowListProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withQuery<TProps, FollowListQuery, FollowListQueryVariables, FollowListProps<TChildProps, TDataName>>(FollowListDocument, {
      alias: 'followList',
      ...operationOptions
    });
};

/**
 * __useFollowListQuery__
 *
 * To run a query within a React component, call `useFollowListQuery` and pass it any options that fit your needs.
 * When your component renders, `useFollowListQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useFollowListQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useFollowListQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<FollowListQuery, FollowListQueryVariables>) {
        return ApolloReactHooks.useQuery<FollowListQuery, FollowListQueryVariables>(FollowListDocument, baseOptions);
      }
export function useFollowListLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<FollowListQuery, FollowListQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<FollowListQuery, FollowListQueryVariables>(FollowListDocument, baseOptions);
        }
export type FollowListQueryHookResult = ReturnType<typeof useFollowListQuery>;
export type FollowListLazyQueryHookResult = ReturnType<typeof useFollowListLazyQuery>;
export type FollowListQueryResult = ApolloReactCommon.QueryResult<FollowListQuery, FollowListQueryVariables>;
export const IsFollowDocument = gql`
    query IsFollow($id: Int!) {
  IsFollow(id: $id)
}
    `;
export type IsFollowComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<IsFollowQuery, IsFollowQueryVariables>, 'query'> & ({ variables: IsFollowQueryVariables; skip?: boolean; } | { skip: boolean; });

    export const IsFollowComponent = (props: IsFollowComponentProps) => (
      <ApolloReactComponents.Query<IsFollowQuery, IsFollowQueryVariables> query={IsFollowDocument} {...props} />
    );
    
export type IsFollowProps<TChildProps = {}, TDataName extends string = 'data'> = {
      [key in TDataName]: ApolloReactHoc.DataValue<IsFollowQuery, IsFollowQueryVariables>
    } & TChildProps;
export function withIsFollow<TProps, TChildProps = {}, TDataName extends string = 'data'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  IsFollowQuery,
  IsFollowQueryVariables,
  IsFollowProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withQuery<TProps, IsFollowQuery, IsFollowQueryVariables, IsFollowProps<TChildProps, TDataName>>(IsFollowDocument, {
      alias: 'isFollow',
      ...operationOptions
    });
};

/**
 * __useIsFollowQuery__
 *
 * To run a query within a React component, call `useIsFollowQuery` and pass it any options that fit your needs.
 * When your component renders, `useIsFollowQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useIsFollowQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useIsFollowQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<IsFollowQuery, IsFollowQueryVariables>) {
        return ApolloReactHooks.useQuery<IsFollowQuery, IsFollowQueryVariables>(IsFollowDocument, baseOptions);
      }
export function useIsFollowLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<IsFollowQuery, IsFollowQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<IsFollowQuery, IsFollowQueryVariables>(IsFollowDocument, baseOptions);
        }
export type IsFollowQueryHookResult = ReturnType<typeof useIsFollowQuery>;
export type IsFollowLazyQueryHookResult = ReturnType<typeof useIsFollowLazyQuery>;
export type IsFollowQueryResult = ApolloReactCommon.QueryResult<IsFollowQuery, IsFollowQueryVariables>;
export const FollowDocument = gql`
    mutation Follow($id: Int!) {
  Follow(id: $id)
}
    `;
export type FollowMutationFn = ApolloReactCommon.MutationFunction<FollowMutation, FollowMutationVariables>;
export type FollowComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<FollowMutation, FollowMutationVariables>, 'mutation'>;

    export const FollowComponent = (props: FollowComponentProps) => (
      <ApolloReactComponents.Mutation<FollowMutation, FollowMutationVariables> mutation={FollowDocument} {...props} />
    );
    
export type FollowProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<FollowMutation, FollowMutationVariables>
    } & TChildProps;
export function withFollow<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  FollowMutation,
  FollowMutationVariables,
  FollowProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, FollowMutation, FollowMutationVariables, FollowProps<TChildProps, TDataName>>(FollowDocument, {
      alias: 'follow',
      ...operationOptions
    });
};

/**
 * __useFollowMutation__
 *
 * To run a mutation, you first call `useFollowMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useFollowMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [followMutation, { data, loading, error }] = useFollowMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useFollowMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<FollowMutation, FollowMutationVariables>) {
        return ApolloReactHooks.useMutation<FollowMutation, FollowMutationVariables>(FollowDocument, baseOptions);
      }
export type FollowMutationHookResult = ReturnType<typeof useFollowMutation>;
export type FollowMutationResult = ApolloReactCommon.MutationResult<FollowMutation>;
export type FollowMutationOptions = ApolloReactCommon.BaseMutationOptions<FollowMutation, FollowMutationVariables>;
export const UnFollowDocument = gql`
    mutation UnFollow($id: Int!) {
  UnFollow(id: $id)
}
    `;
export type UnFollowMutationFn = ApolloReactCommon.MutationFunction<UnFollowMutation, UnFollowMutationVariables>;
export type UnFollowComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<UnFollowMutation, UnFollowMutationVariables>, 'mutation'>;

    export const UnFollowComponent = (props: UnFollowComponentProps) => (
      <ApolloReactComponents.Mutation<UnFollowMutation, UnFollowMutationVariables> mutation={UnFollowDocument} {...props} />
    );
    
export type UnFollowProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<UnFollowMutation, UnFollowMutationVariables>
    } & TChildProps;
export function withUnFollow<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  UnFollowMutation,
  UnFollowMutationVariables,
  UnFollowProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, UnFollowMutation, UnFollowMutationVariables, UnFollowProps<TChildProps, TDataName>>(UnFollowDocument, {
      alias: 'unFollow',
      ...operationOptions
    });
};

/**
 * __useUnFollowMutation__
 *
 * To run a mutation, you first call `useUnFollowMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUnFollowMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [unFollowMutation, { data, loading, error }] = useUnFollowMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useUnFollowMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<UnFollowMutation, UnFollowMutationVariables>) {
        return ApolloReactHooks.useMutation<UnFollowMutation, UnFollowMutationVariables>(UnFollowDocument, baseOptions);
      }
export type UnFollowMutationHookResult = ReturnType<typeof useUnFollowMutation>;
export type UnFollowMutationResult = ApolloReactCommon.MutationResult<UnFollowMutation>;
export type UnFollowMutationOptions = ApolloReactCommon.BaseMutationOptions<UnFollowMutation, UnFollowMutationVariables>;
export const UsersDocument = gql`
    query Users($username: String = null) {
  Users(username: $username) {
    edges {
      node {
        id
        username
        avatar
        FansNum
        FollowNum
        LikeNum
        Words
      }
    }
    pageInfo {
      endCursor
      hasNextPage
    }
  }
}
    `;
export type UsersComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<UsersQuery, UsersQueryVariables>, 'query'>;

    export const UsersComponent = (props: UsersComponentProps) => (
      <ApolloReactComponents.Query<UsersQuery, UsersQueryVariables> query={UsersDocument} {...props} />
    );
    
export type UsersProps<TChildProps = {}, TDataName extends string = 'data'> = {
      [key in TDataName]: ApolloReactHoc.DataValue<UsersQuery, UsersQueryVariables>
    } & TChildProps;
export function withUsers<TProps, TChildProps = {}, TDataName extends string = 'data'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  UsersQuery,
  UsersQueryVariables,
  UsersProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withQuery<TProps, UsersQuery, UsersQueryVariables, UsersProps<TChildProps, TDataName>>(UsersDocument, {
      alias: 'users',
      ...operationOptions
    });
};

/**
 * __useUsersQuery__
 *
 * To run a query within a React component, call `useUsersQuery` and pass it any options that fit your needs.
 * When your component renders, `useUsersQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useUsersQuery({
 *   variables: {
 *      username: // value for 'username'
 *   },
 * });
 */
export function useUsersQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<UsersQuery, UsersQueryVariables>) {
        return ApolloReactHooks.useQuery<UsersQuery, UsersQueryVariables>(UsersDocument, baseOptions);
      }
export function useUsersLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<UsersQuery, UsersQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<UsersQuery, UsersQueryVariables>(UsersDocument, baseOptions);
        }
export type UsersQueryHookResult = ReturnType<typeof useUsersQuery>;
export type UsersLazyQueryHookResult = ReturnType<typeof useUsersLazyQuery>;
export type UsersQueryResult = ApolloReactCommon.QueryResult<UsersQuery, UsersQueryVariables>;
export const ViewAddDocument = gql`
    mutation ViewAdd($id: Int!) {
  ViewAdd(id: $id)
}
    `;
export type ViewAddMutationFn = ApolloReactCommon.MutationFunction<ViewAddMutation, ViewAddMutationVariables>;
export type ViewAddComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<ViewAddMutation, ViewAddMutationVariables>, 'mutation'>;

    export const ViewAddComponent = (props: ViewAddComponentProps) => (
      <ApolloReactComponents.Mutation<ViewAddMutation, ViewAddMutationVariables> mutation={ViewAddDocument} {...props} />
    );
    
export type ViewAddProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<ViewAddMutation, ViewAddMutationVariables>
    } & TChildProps;
export function withViewAdd<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  ViewAddMutation,
  ViewAddMutationVariables,
  ViewAddProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, ViewAddMutation, ViewAddMutationVariables, ViewAddProps<TChildProps, TDataName>>(ViewAddDocument, {
      alias: 'viewAdd',
      ...operationOptions
    });
};

/**
 * __useViewAddMutation__
 *
 * To run a mutation, you first call `useViewAddMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useViewAddMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [viewAddMutation, { data, loading, error }] = useViewAddMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useViewAddMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<ViewAddMutation, ViewAddMutationVariables>) {
        return ApolloReactHooks.useMutation<ViewAddMutation, ViewAddMutationVariables>(ViewAddDocument, baseOptions);
      }
export type ViewAddMutationHookResult = ReturnType<typeof useViewAddMutation>;
export type ViewAddMutationResult = ApolloReactCommon.MutationResult<ViewAddMutation>;
export type ViewAddMutationOptions = ApolloReactCommon.BaseMutationOptions<ViewAddMutation, ViewAddMutationVariables>;
export const LikeDocument = gql`
    mutation Like($id: Int!, $objtyp: ObjType!) {
  Like(id: $id, objType: $objtyp)
}
    `;
export type LikeMutationFn = ApolloReactCommon.MutationFunction<LikeMutation, LikeMutationVariables>;
export type LikeComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<LikeMutation, LikeMutationVariables>, 'mutation'>;

    export const LikeComponent = (props: LikeComponentProps) => (
      <ApolloReactComponents.Mutation<LikeMutation, LikeMutationVariables> mutation={LikeDocument} {...props} />
    );
    
export type LikeProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<LikeMutation, LikeMutationVariables>
    } & TChildProps;
export function withLike<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  LikeMutation,
  LikeMutationVariables,
  LikeProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, LikeMutation, LikeMutationVariables, LikeProps<TChildProps, TDataName>>(LikeDocument, {
      alias: 'like',
      ...operationOptions
    });
};

/**
 * __useLikeMutation__
 *
 * To run a mutation, you first call `useLikeMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useLikeMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [likeMutation, { data, loading, error }] = useLikeMutation({
 *   variables: {
 *      id: // value for 'id'
 *      objtyp: // value for 'objtyp'
 *   },
 * });
 */
export function useLikeMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<LikeMutation, LikeMutationVariables>) {
        return ApolloReactHooks.useMutation<LikeMutation, LikeMutationVariables>(LikeDocument, baseOptions);
      }
export type LikeMutationHookResult = ReturnType<typeof useLikeMutation>;
export type LikeMutationResult = ApolloReactCommon.MutationResult<LikeMutation>;
export type LikeMutationOptions = ApolloReactCommon.BaseMutationOptions<LikeMutation, LikeMutationVariables>;
export const UnLikeDocument = gql`
    mutation UnLike($id: Int!, $objtyp: ObjType!) {
  Unlike(id: $id, objType: $objtyp)
}
    `;
export type UnLikeMutationFn = ApolloReactCommon.MutationFunction<UnLikeMutation, UnLikeMutationVariables>;
export type UnLikeComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<UnLikeMutation, UnLikeMutationVariables>, 'mutation'>;

    export const UnLikeComponent = (props: UnLikeComponentProps) => (
      <ApolloReactComponents.Mutation<UnLikeMutation, UnLikeMutationVariables> mutation={UnLikeDocument} {...props} />
    );
    
export type UnLikeProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<UnLikeMutation, UnLikeMutationVariables>
    } & TChildProps;
export function withUnLike<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  UnLikeMutation,
  UnLikeMutationVariables,
  UnLikeProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, UnLikeMutation, UnLikeMutationVariables, UnLikeProps<TChildProps, TDataName>>(UnLikeDocument, {
      alias: 'unLike',
      ...operationOptions
    });
};

/**
 * __useUnLikeMutation__
 *
 * To run a mutation, you first call `useUnLikeMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUnLikeMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [unLikeMutation, { data, loading, error }] = useUnLikeMutation({
 *   variables: {
 *      id: // value for 'id'
 *      objtyp: // value for 'objtyp'
 *   },
 * });
 */
export function useUnLikeMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<UnLikeMutation, UnLikeMutationVariables>) {
        return ApolloReactHooks.useMutation<UnLikeMutation, UnLikeMutationVariables>(UnLikeDocument, baseOptions);
      }
export type UnLikeMutationHookResult = ReturnType<typeof useUnLikeMutation>;
export type UnLikeMutationResult = ApolloReactCommon.MutationResult<UnLikeMutation>;
export type UnLikeMutationOptions = ApolloReactCommon.BaseMutationOptions<UnLikeMutation, UnLikeMutationVariables>;
export const HasLikeDocument = gql`
    query HasLike($id: Int!, $objtyp: ObjType!) {
  HasLike(id: $id, objType: $objtyp)
}
    `;
export type HasLikeComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<HasLikeQuery, HasLikeQueryVariables>, 'query'> & ({ variables: HasLikeQueryVariables; skip?: boolean; } | { skip: boolean; });

    export const HasLikeComponent = (props: HasLikeComponentProps) => (
      <ApolloReactComponents.Query<HasLikeQuery, HasLikeQueryVariables> query={HasLikeDocument} {...props} />
    );
    
export type HasLikeProps<TChildProps = {}, TDataName extends string = 'data'> = {
      [key in TDataName]: ApolloReactHoc.DataValue<HasLikeQuery, HasLikeQueryVariables>
    } & TChildProps;
export function withHasLike<TProps, TChildProps = {}, TDataName extends string = 'data'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  HasLikeQuery,
  HasLikeQueryVariables,
  HasLikeProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withQuery<TProps, HasLikeQuery, HasLikeQueryVariables, HasLikeProps<TChildProps, TDataName>>(HasLikeDocument, {
      alias: 'hasLike',
      ...operationOptions
    });
};

/**
 * __useHasLikeQuery__
 *
 * To run a query within a React component, call `useHasLikeQuery` and pass it any options that fit your needs.
 * When your component renders, `useHasLikeQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useHasLikeQuery({
 *   variables: {
 *      id: // value for 'id'
 *      objtyp: // value for 'objtyp'
 *   },
 * });
 */
export function useHasLikeQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<HasLikeQuery, HasLikeQueryVariables>) {
        return ApolloReactHooks.useQuery<HasLikeQuery, HasLikeQueryVariables>(HasLikeDocument, baseOptions);
      }
export function useHasLikeLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<HasLikeQuery, HasLikeQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<HasLikeQuery, HasLikeQueryVariables>(HasLikeDocument, baseOptions);
        }
export type HasLikeQueryHookResult = ReturnType<typeof useHasLikeQuery>;
export type HasLikeLazyQueryHookResult = ReturnType<typeof useHasLikeLazyQuery>;
export type HasLikeQueryResult = ApolloReactCommon.QueryResult<HasLikeQuery, HasLikeQueryVariables>;
export const ReplyListDocument = gql`
    query ReplyList($id: Int!) {
  ReplyList(id: $id) {
    id
    content
    updatedAt
    User {
      id
      username
      avatar
    }
  }
}
    `;
export type ReplyListComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<ReplyListQuery, ReplyListQueryVariables>, 'query'> & ({ variables: ReplyListQueryVariables; skip?: boolean; } | { skip: boolean; });

    export const ReplyListComponent = (props: ReplyListComponentProps) => (
      <ApolloReactComponents.Query<ReplyListQuery, ReplyListQueryVariables> query={ReplyListDocument} {...props} />
    );
    
export type ReplyListProps<TChildProps = {}, TDataName extends string = 'data'> = {
      [key in TDataName]: ApolloReactHoc.DataValue<ReplyListQuery, ReplyListQueryVariables>
    } & TChildProps;
export function withReplyList<TProps, TChildProps = {}, TDataName extends string = 'data'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  ReplyListQuery,
  ReplyListQueryVariables,
  ReplyListProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withQuery<TProps, ReplyListQuery, ReplyListQueryVariables, ReplyListProps<TChildProps, TDataName>>(ReplyListDocument, {
      alias: 'replyList',
      ...operationOptions
    });
};

/**
 * __useReplyListQuery__
 *
 * To run a query within a React component, call `useReplyListQuery` and pass it any options that fit your needs.
 * When your component renders, `useReplyListQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useReplyListQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useReplyListQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<ReplyListQuery, ReplyListQueryVariables>) {
        return ApolloReactHooks.useQuery<ReplyListQuery, ReplyListQueryVariables>(ReplyListDocument, baseOptions);
      }
export function useReplyListLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<ReplyListQuery, ReplyListQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<ReplyListQuery, ReplyListQueryVariables>(ReplyListDocument, baseOptions);
        }
export type ReplyListQueryHookResult = ReturnType<typeof useReplyListQuery>;
export type ReplyListLazyQueryHookResult = ReturnType<typeof useReplyListLazyQuery>;
export type ReplyListQueryResult = ApolloReactCommon.QueryResult<ReplyListQuery, ReplyListQueryVariables>;
export const AddCommentDocument = gql`
    mutation AddComment($id: Int!, $content: String!) {
  AddComment(id: $id, content: $content) {
    User {
      id
      username
      avatar
    }
    content
    floor
    id
    likeNum
    updatedAt
  }
}
    `;
export type AddCommentMutationFn = ApolloReactCommon.MutationFunction<AddCommentMutation, AddCommentMutationVariables>;
export type AddCommentComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<AddCommentMutation, AddCommentMutationVariables>, 'mutation'>;

    export const AddCommentComponent = (props: AddCommentComponentProps) => (
      <ApolloReactComponents.Mutation<AddCommentMutation, AddCommentMutationVariables> mutation={AddCommentDocument} {...props} />
    );
    
export type AddCommentProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<AddCommentMutation, AddCommentMutationVariables>
    } & TChildProps;
export function withAddComment<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  AddCommentMutation,
  AddCommentMutationVariables,
  AddCommentProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, AddCommentMutation, AddCommentMutationVariables, AddCommentProps<TChildProps, TDataName>>(AddCommentDocument, {
      alias: 'addComment',
      ...operationOptions
    });
};

/**
 * __useAddCommentMutation__
 *
 * To run a mutation, you first call `useAddCommentMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAddCommentMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [addCommentMutation, { data, loading, error }] = useAddCommentMutation({
 *   variables: {
 *      id: // value for 'id'
 *      content: // value for 'content'
 *   },
 * });
 */
export function useAddCommentMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<AddCommentMutation, AddCommentMutationVariables>) {
        return ApolloReactHooks.useMutation<AddCommentMutation, AddCommentMutationVariables>(AddCommentDocument, baseOptions);
      }
export type AddCommentMutationHookResult = ReturnType<typeof useAddCommentMutation>;
export type AddCommentMutationResult = ApolloReactCommon.MutationResult<AddCommentMutation>;
export type AddCommentMutationOptions = ApolloReactCommon.BaseMutationOptions<AddCommentMutation, AddCommentMutationVariables>;
export const AddReplyDocument = gql`
    mutation AddReply($id: Int!, $content: String!) {
  AddReply(id: $id, content: $content) {
    User {
      id
      username
      avatar
    }
    content
    id
    updatedAt
  }
}
    `;
export type AddReplyMutationFn = ApolloReactCommon.MutationFunction<AddReplyMutation, AddReplyMutationVariables>;
export type AddReplyComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<AddReplyMutation, AddReplyMutationVariables>, 'mutation'>;

    export const AddReplyComponent = (props: AddReplyComponentProps) => (
      <ApolloReactComponents.Mutation<AddReplyMutation, AddReplyMutationVariables> mutation={AddReplyDocument} {...props} />
    );
    
export type AddReplyProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<AddReplyMutation, AddReplyMutationVariables>
    } & TChildProps;
export function withAddReply<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  AddReplyMutation,
  AddReplyMutationVariables,
  AddReplyProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, AddReplyMutation, AddReplyMutationVariables, AddReplyProps<TChildProps, TDataName>>(AddReplyDocument, {
      alias: 'addReply',
      ...operationOptions
    });
};

/**
 * __useAddReplyMutation__
 *
 * To run a mutation, you first call `useAddReplyMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAddReplyMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [addReplyMutation, { data, loading, error }] = useAddReplyMutation({
 *   variables: {
 *      id: // value for 'id'
 *      content: // value for 'content'
 *   },
 * });
 */
export function useAddReplyMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<AddReplyMutation, AddReplyMutationVariables>) {
        return ApolloReactHooks.useMutation<AddReplyMutation, AddReplyMutationVariables>(AddReplyDocument, baseOptions);
      }
export type AddReplyMutationHookResult = ReturnType<typeof useAddReplyMutation>;
export type AddReplyMutationResult = ApolloReactCommon.MutationResult<AddReplyMutation>;
export type AddReplyMutationOptions = ApolloReactCommon.BaseMutationOptions<AddReplyMutation, AddReplyMutationVariables>;
export const AddMsgDocument = gql`
    mutation AddMsg($typ: MsgType!, $fromId: Int!, $toId: Int!, $content: String!) {
  AddMsg(typ: $typ, fromId: $fromId, toId: $toId, content: $content)
}
    `;
export type AddMsgMutationFn = ApolloReactCommon.MutationFunction<AddMsgMutation, AddMsgMutationVariables>;
export type AddMsgComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<AddMsgMutation, AddMsgMutationVariables>, 'mutation'>;

    export const AddMsgComponent = (props: AddMsgComponentProps) => (
      <ApolloReactComponents.Mutation<AddMsgMutation, AddMsgMutationVariables> mutation={AddMsgDocument} {...props} />
    );
    
export type AddMsgProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<AddMsgMutation, AddMsgMutationVariables>
    } & TChildProps;
export function withAddMsg<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  AddMsgMutation,
  AddMsgMutationVariables,
  AddMsgProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, AddMsgMutation, AddMsgMutationVariables, AddMsgProps<TChildProps, TDataName>>(AddMsgDocument, {
      alias: 'addMsg',
      ...operationOptions
    });
};

/**
 * __useAddMsgMutation__
 *
 * To run a mutation, you first call `useAddMsgMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAddMsgMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [addMsgMutation, { data, loading, error }] = useAddMsgMutation({
 *   variables: {
 *      typ: // value for 'typ'
 *      fromId: // value for 'fromId'
 *      toId: // value for 'toId'
 *      content: // value for 'content'
 *   },
 * });
 */
export function useAddMsgMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<AddMsgMutation, AddMsgMutationVariables>) {
        return ApolloReactHooks.useMutation<AddMsgMutation, AddMsgMutationVariables>(AddMsgDocument, baseOptions);
      }
export type AddMsgMutationHookResult = ReturnType<typeof useAddMsgMutation>;
export type AddMsgMutationResult = ApolloReactCommon.MutationResult<AddMsgMutation>;
export type AddMsgMutationOptions = ApolloReactCommon.BaseMutationOptions<AddMsgMutation, AddMsgMutationVariables>;
export const MsgNumDocument = gql`
    query MsgNum {
  MsgNum {
    comment
    follow
    like
    reply
  }
}
    `;
export type MsgNumComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<MsgNumQuery, MsgNumQueryVariables>, 'query'>;

    export const MsgNumComponent = (props: MsgNumComponentProps) => (
      <ApolloReactComponents.Query<MsgNumQuery, MsgNumQueryVariables> query={MsgNumDocument} {...props} />
    );
    
export type MsgNumProps<TChildProps = {}, TDataName extends string = 'data'> = {
      [key in TDataName]: ApolloReactHoc.DataValue<MsgNumQuery, MsgNumQueryVariables>
    } & TChildProps;
export function withMsgNum<TProps, TChildProps = {}, TDataName extends string = 'data'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  MsgNumQuery,
  MsgNumQueryVariables,
  MsgNumProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withQuery<TProps, MsgNumQuery, MsgNumQueryVariables, MsgNumProps<TChildProps, TDataName>>(MsgNumDocument, {
      alias: 'msgNum',
      ...operationOptions
    });
};

/**
 * __useMsgNumQuery__
 *
 * To run a query within a React component, call `useMsgNumQuery` and pass it any options that fit your needs.
 * When your component renders, `useMsgNumQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useMsgNumQuery({
 *   variables: {
 *   },
 * });
 */
export function useMsgNumQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<MsgNumQuery, MsgNumQueryVariables>) {
        return ApolloReactHooks.useQuery<MsgNumQuery, MsgNumQueryVariables>(MsgNumDocument, baseOptions);
      }
export function useMsgNumLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<MsgNumQuery, MsgNumQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<MsgNumQuery, MsgNumQueryVariables>(MsgNumDocument, baseOptions);
        }
export type MsgNumQueryHookResult = ReturnType<typeof useMsgNumQuery>;
export type MsgNumLazyQueryHookResult = ReturnType<typeof useMsgNumLazyQuery>;
export type MsgNumQueryResult = ApolloReactCommon.QueryResult<MsgNumQuery, MsgNumQueryVariables>;
export const ListMsgDocument = gql`
    query ListMsg($typ: MsgType!) {
  ListMsg(typ: $typ) {
    User {
      id
      username
      avatar
    }
    content
    updatedAt
  }
}
    `;
export type ListMsgComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<ListMsgQuery, ListMsgQueryVariables>, 'query'> & ({ variables: ListMsgQueryVariables; skip?: boolean; } | { skip: boolean; });

    export const ListMsgComponent = (props: ListMsgComponentProps) => (
      <ApolloReactComponents.Query<ListMsgQuery, ListMsgQueryVariables> query={ListMsgDocument} {...props} />
    );
    
export type ListMsgProps<TChildProps = {}, TDataName extends string = 'data'> = {
      [key in TDataName]: ApolloReactHoc.DataValue<ListMsgQuery, ListMsgQueryVariables>
    } & TChildProps;
export function withListMsg<TProps, TChildProps = {}, TDataName extends string = 'data'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  ListMsgQuery,
  ListMsgQueryVariables,
  ListMsgProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withQuery<TProps, ListMsgQuery, ListMsgQueryVariables, ListMsgProps<TChildProps, TDataName>>(ListMsgDocument, {
      alias: 'listMsg',
      ...operationOptions
    });
};

/**
 * __useListMsgQuery__
 *
 * To run a query within a React component, call `useListMsgQuery` and pass it any options that fit your needs.
 * When your component renders, `useListMsgQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useListMsgQuery({
 *   variables: {
 *      typ: // value for 'typ'
 *   },
 * });
 */
export function useListMsgQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<ListMsgQuery, ListMsgQueryVariables>) {
        return ApolloReactHooks.useQuery<ListMsgQuery, ListMsgQueryVariables>(ListMsgDocument, baseOptions);
      }
export function useListMsgLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<ListMsgQuery, ListMsgQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<ListMsgQuery, ListMsgQueryVariables>(ListMsgDocument, baseOptions);
        }
export type ListMsgQueryHookResult = ReturnType<typeof useListMsgQuery>;
export type ListMsgLazyQueryHookResult = ReturnType<typeof useListMsgLazyQuery>;
export type ListMsgQueryResult = ApolloReactCommon.QueryResult<ListMsgQuery, ListMsgQueryVariables>;