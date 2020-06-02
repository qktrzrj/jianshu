import React, {Suspense, lazy} from "react";
import {BrowserRouter, Route, Switch} from "react-router-dom"
import Loading from "./component/loading";
import Layout from "./component/layout/BasicLayout";

const Index = lazy(() => import('./pages/index'))
const UserCenter = lazy(() => import("./pages/userCenter"))
const Setting = lazy(() => import('./pages/setting'))
const Subscriptions = lazy(() => import('./pages/subscriptions'))


export default function App() {
    return (
        <BrowserRouter>
            <Suspense fallback={<Loading/>}>
                <Switch>
                    <Route path='/signIn' component={lazy(() => import('./pages/sign'))}/>
                    <Route path='/signUp' component={lazy(() => import('./pages/sign'))}/>
                    <Route path='/writer' component={lazy(() => import('./pages/writer'))}/>
                    <Route render={(props => (
                        <Layout {...props} render={currentUser => (
                            <React.Fragment>
                                <Switch>
                                    <Route path='//' children={<Index/>}/>
                                    <Route path='/u/:id' render={p => <UserCenter {...p} currentUser={currentUser}/>}/>
                                    {currentUser &&
                                    <Route path='/setting/:key'
                                           render={p => <Setting {...p} currentUser={currentUser}/>}/>
                                    }
                                    {currentUser &&
                                    <Route path='/subscriptions'
                                           render={p => <Subscriptions {...p} currentUser={currentUser}/>}/>
                                    }
                                    <Route component={lazy(() => import('./404'))}/>
                                </Switch>
                            </React.Fragment>
                        )}/>
                    ))}/>
                </Switch>
            </Suspense>
        </BrowserRouter>
    )
}

