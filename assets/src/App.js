import { Switch, Route, Redirect } from 'react-router-dom';
import { CompatRouter } from "react-router-dom-v5-compat";

import Layout from './components/Layout/Layout';
import HomePage from './pages/HomePage';
import AuthPage from './pages/AuthPage';
import PostsPage from './pages/PostsPage';
import { useContext } from 'react';
import AuthContext from './store/auth-context';

function App() {
  const authContext = useContext(AuthContext);

  return (
    <Layout>
      <CompatRouter>
        <Switch>
          <Route path="/" exact>
            <HomePage/>
          </Route>
          {!authContext.loggedIn && (
            <Route path="/auth">
              <AuthPage />
            </Route>
          )}
          <Route path="/posts">
            {authContext.loggedIn && <PostsPage />}
            {!authContext.loggedIn && <Redirect to="/auth" />}
          </Route>
          <Route path="*">
            <Redirect to="/"/>
          </Route>
        </Switch>
      </CompatRouter>
    </Layout>
  );
}

export default App;
