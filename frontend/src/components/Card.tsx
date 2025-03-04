import React from "react"
import { ReactComponent as Bookmark} from "../images/bookmark.svg"
import {useHistory } from "react-router-dom"

function Card(props:any) {

  const history = useHistory()
  const { card, bookmarkSet } = props;

  function addBookmarkHandeler(){
    let data = {
      Title : card.name,
      Url : card.website
    }
    
    const requestOptions: RequestInit = {
        mode : 'no-cors',
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data),
    };
    fetch('http://localhost:9000/create',requestOptions).then((res) => history.push('/bookmarks'))
  }
  const imgsrc = card.imageurl
  return (
    <article className='article'>
      {/* <a href={card.website} target="_blank" className='cardlink'></a> */}
      <div className='cardtop'>
        <div className='cardheader'>
          <a href={card.website} target="_blank" rel="noreferrer" style={{textDecoration: "inherit", color: "inherit"}}><h3>{card.name}</h3></a>
        </div>
        <div className="description"><p>{card.description}</p></div>
      </div>
      <div className='cardmiddle'>
        <img src={imgsrc} alt={card?.title}/>
      </div> 
      <div className='icon-1' style={{display : "flex",justifyContent: "flex-end", cursor: 'pointer'}}>
      <span  onClick={(e) => {addBookmarkHandeler()}}>{!bookmarkSet && <Bookmark/>}</span>
        </div>    
    </article>   
  )
}

export default Card