import React, {Suspense, lazy} from "react";
import {BrowserRouter, Route, Switch} from "react-router-dom"
import Loading from "./component/loading";

export default function App() {
    return (
        <BrowserRouter>
            <Suspense fallback={<Loading/>}>
                <Switch>
                    <Route path='/signIn' component={lazy(() => import('./pages/sign/sign'))}/>
                    <Route path='/signUp' component={lazy(() => import('./pages/sign/sign'))}/>
                    <Route path='/' component={lazy(() => import('./pages/jianshu/jianshu'))}/>
                </Switch>
            </Suspense>
        </BrowserRouter>
    )
}
