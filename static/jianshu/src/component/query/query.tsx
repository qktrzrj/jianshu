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
     updatedAt
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
    subTitle
    content
    state
    updatedAt
    ViewNum
    LikeNum
    CmtNum
    User{
        ...userInfo
     }
     CommentList{
        id
        floor
        content
        updatedAt
        likeNum
        User{
          id
          username
          avatar
        }
      }
  }
}
${UserInfoFragmentGQL}
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
query Articles($cursor:String=null,$uid:Int=null,$condition:String=null){
  Articles(first:10,after:$cursor,uid:$uid,condition:$condition){
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

export const LikeArticlesGQL = gql`
query LikeArticles($cursor:String=null){
  CurLikeArticles(first:10,after:$cursor){
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

export const FollowGQL = gql`
mutation Follow($id:Int!){
  Follow(id:$id)
}
`

export const UnFollowGQL = gql`
mutation UnFollow($id:Int!){
  UnFollow(id:$id)
}
`

export const UsersGQL = gql`
query Users($username:String=null){
    Users(username:$username){
        edges{
          node{
            id
            username
            avatar
            FansNum
            FollowNum
            LikeNum
            Words
          }
        }
        pageInfo{
          endCursor
          hasNextPage
        }
    }
}
`

export const ViewGQL=gql`
mutation ViewAdd($id:Int!){
  ViewAdd(id:$id)
}
`

export const LikeGQL=gql`
mutation Like($id:Int!,$objtyp:ObjType!){
  Like(id:$id,objType:$objtyp)
}
`

export const UnLikeGQL=gql`
mutation UnLike($id:Int!,$objtyp:ObjType!){
  Unlike(id:$id,objType:$objtyp)
}
`

export const HasLikeGQL=gql`
query HasLike($id:Int!,$objtyp:ObjType!){
  HasLike(id:$id,objType:$objtyp)
}
`

export const ReplyListGQL =gql`
query ReplyList($id:Int!){
    ReplyList(id:$id){
            id
            content
            updatedAt
            User{
                id
                username
                avatar
            }
     }
}
`

export const AddCommentGQL = gql`
mutation AddComment($id:Int!,$content:String!){
  AddComment(id:$id,content:$content){
    User{
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
`

export const AddReplyGQL=gql`
mutation AddReply($id:Int!,$content:String!){
  AddReply(id:$id,content:$content){
    User{
      id
      username
      avatar
    }
    content
    id
    updatedAt
  }
}
`

export const AddMsgGQl=gql`
mutation AddMsg($typ:MsgType!,$fromId:Int!,$toId:Int!,$content:String!){
  AddMsg(typ:$typ,fromId:$fromId,toId:$toId,content:$content)
}
`

export const MsgNumGQl=gql`
query MsgNum{
  MsgNum{
    comment
    follow
    like
    reply
  }
}
`

export const ListMsgGQL=gql`
query ListMsg($typ:MsgType!){
  ListMsg(typ:$typ){
    User{
      id
      username
      avatar
    }
    content
    updatedAt
  }
}
`

