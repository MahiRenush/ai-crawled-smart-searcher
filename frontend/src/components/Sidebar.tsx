import React from 'react'
import {
    Link,
    Switch,
    Route
  } from "react-router-dom";
import { ReactComponent as Popular } from "../images/popular.svg"
import { ReactComponent as Upvoted } from "../images/mostupvoted.svg"
import { ReactComponent as Search } from "../images/search.svg"
import { ReactComponent as Bookmark} from "../images/bookmark.svg"
import { useLocation } from 'react-router-dom'

function Sidebar({routes}: {routes: any}) {
  let location = useLocation();

  let isPopular = location.pathname.includes("/popular")
  let isVoted = location.pathname.includes("/upvoted")
  let isSearch = location.pathname.includes("/search")
  let isBookmark = location.pathname.includes("/bookmarks")


  return (
    <div className="sidebar bg-color">
          <ul className='sidebarlist'>                        
            <li className='listheader'>Explore</li>
            <li className={isSearch ? 'list active': 'list'}><Link to='/search'><span className='icon'><Search/></span> <span className='listtext'>Search</span></Link></li>
            <li className={isPopular ? 'list active': 'list'}><Link to='/popular'><span className='icon'><Popular/></span> <span className='listtext'>Popular</span></Link></li>
            <li className={isVoted ? 'list active': 'list'}><Link to='/upvoted'><span className='icon'><Upvoted/></span> <span className='listtext'>Most upvoted</span></Link></li>
            <li className={isBookmark ? 'list active': 'list'}><Link to='/bookmarks'><span className='icon'><Bookmark/></span> <span className='listtext'>Bookmarks</span></Link></li>
          </ul>

          <Switch>
            {routes.map((route: any, index: number) => (
              <Route
                key={index}
                path={route.path}
                exact={route.exact}
              />
            ))}
          </Switch>
        </div>
  )
}

export default Sidebar