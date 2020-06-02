import {gql} from "apollo-boost";

export const CheckUsernameGQL = gql`
query ValidUsername($username:String!){
  ValidUsername(username:$username)
}
`;

export const CheckEmailGQL = gql`
query ValidEmail($email:String!){
  ValidEmail(email:$email)
}
`;

export const SignUpGQL = gql`
mutation SignUp($email:String!,$password:String!,$username:String!){
  SignUp(email:$email,password:$password,username:$username){
    id
  }
}
`;

export const SignInGQL = gql`
mutation SignIn($username:String!,$password:String!,$rememberme:Boolean!){
  SignIn(username:$username,password:$password,rememberme:$rememberme){
    id
  }
}
`

export const LogoutGQL = gql`
mutation Logout{
  Logout
}
`

const UserInfoFragmentGQL = gql`
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
`

export const CurrentUserGQL = gql`
query CurrentUser{
  CurrentUser{
    ...userInfo
  }
}
${UserInfoFragmentGQL}
`

const ArticlesFragmentGQL = gql`
fragment articlesInfo on Article {
     id
     title
     subTitle
     cover
     content
     User{
        id
        username
     }
     ViewNum
     LikeNum
     CmtNum
}
`

export const HotArticlesGQL = gql`
query HotArticles($cursor:String){
  HotArticles(first:10,after:$cursor){
    edges{
      node{
        ...articlesInfo
      }
    }
    pageInfo{
      endCursor
      hasNextPage
    }
  }
}
${ArticlesFragmentGQL}
`

export const DraftArticleGQL = gql`
mutation DraftArticle($title:String!){
  DraftArticle(title:$title){
    id
    title
    state
  }
}
`

export const MyArticles = gql`
query MyArticles{
  CurArticles{
    edges{
      node{
        id
        title
        state
      }
    }
  }
}
`

export const ArticleGQL = gql`
query Article($id:Int!){
  Article(id:$id){
    id
    title
    content
    state
  }
}
`

export const DeleteArticle = gql`
mutation DeleteArticle($id:Int!){
  DeleteArticle(id:$id)
}
`

export const UpdateArticleGQL = gql`
mutation UpdateArticle($id:Int!,$content:String,$cover:String,$subTitle:String,$title:String){
  UpdateArticle(id:$id,content:$content,cover:$cover,subTitle:$subTitle,title:$title){
    id
    title
    content
    state
  }
}
`

export const NewArticleGQL = gql`
mutation NewArticle($id:Int!){
  NewArticle(id:$id){
    id
    state
  }
}
`

export const UploadGQL = gql`
mutation Upload($file:Upload!){
  Upload(file:$file)
}
`

export const UserGQL = gql`
query User($id:Int!){
  User(id:$id){
    ...userInfo
  }
}
${UserInfoFragmentGQL}
`

export const UpdateUserInfoGQL = gql`
mutation UpdateUserInfo($username:String=null,$avatar:String=null,$email:String=null,$gender:Gender=null,$introduce:String=null,$password:String=null){
 UpdateUserInfo(username:$username,avatar:$avatar,email:$email,gender:$gender,introduce:$introduce,password:$password)
}
`

export const ArticlesGQL = gql`
query Articles($cursor:String,$uid:Int){
  Articles(first:10,after:$cursor,uid:$uid){
    edges{
      node{
        ...articlesInfo
      }
    }
    pageInfo{
      endCursor
      hasNextPage
    }
  }
}
`

export const FollowListGQL = gql`
query FollowList($id:Int!) {
    Followed(id:$id){
      ...userInfo
    }
  }
${UserInfoFragmentGQL}
`

export const IsFollowGQL = gql`
query IsFollow($id:Int!) {
  IsFollow(id:$id)
}
`

export const FollowGQL =gql`
mutation Follow($id:Int!){
  Follow(id:$id)
}
`

export const UnFollowGQL =gql`
mutation UnFollow($id:Int!){
  UnFollow(id:$id)
}
`