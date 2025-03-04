import { Switch, Route } from "react-router-dom";

function Main({ routes }: { routes: any }) {
  return (
    <div className="main">
      <Switch>
        {routes.map((route: any, index: number) => (
          <Route
            key={index}
            path={route.path}
            exact={route.exact}
            children={<route.component />}
          />
        ))}
      </Switch>
    </div>
  );
}

export default Main;