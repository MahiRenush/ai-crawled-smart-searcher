import { BrowserRouter as Router } from "react-router-dom";
import Header from "./components/Header";
import Sidebar from "./components/Sidebar";
import Main from "./components/Main";
import Popular from "./components/Popular"
import MostUpvoted from "./components/MostUpvoted"
import Search from "./components/Search"
import "./App.css"
import Bookmarks from "./components/Bookmarks";

const routes = [
  {
    path: "/",
    exact: true,
    component: () => <Search />
  },
  {
    path: "/search",
    component: () => <Search />
  },
  {
    path: "/popular",
    component: () => <Popular />
  },
  {
    path: "/upvoted",
    component: () => <MostUpvoted />
  },
  {
    path: "/bookmarks",
    component: () => <Bookmarks />
  }
];

export default function App() {
  return (
    <Router>
      <Header />
      <div className="body bg-color">
        <Sidebar routes={routes}/>
        <Main routes={routes} />
      </div>
    </Router>
  );
}
