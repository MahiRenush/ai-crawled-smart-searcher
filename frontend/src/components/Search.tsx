import {KeyboardEvent, useState, useEffect} from "react"
import Card from './Card'
import { ReactComponent as SearchSvg } from "../images/search.svg"
import { ReactComponent as Clear } from "../images/clear.svg"


function Search() {
  const [searchText, setSearchText] =  useState<string>("");
  const [articles, setArticles] =  useState<any>([]);
  const [bookmarks,setbookmarks] = useState<any>([])
  


  var clearClass = `${searchText ? "clearButton" : "invisible" }`

  const handleChange = (e: any) => {
    let value  = e.target.value;
    setSearchText(value)
    if (value === "") {
      fetch(`http://localhost:9000?q=`, {method: 'get', mode: 'cors'})
      .then((res) => res.json())
      .then((data) => setArticles(data))
    }
  }

  const clearClick = () => {
    setSearchText("")
    fetch(`http://localhost:9000?q=`, {method: 'get', mode: 'cors'})
      .then((res) => res.json())
      .then((data) => setArticles(data))
  }

  const searchLinks = (e: KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter' || e.keyCode === 13) {     
      fetch(`http://localhost:9000?q=${searchText}`, {method: 'GET', mode: 'cors'})
      .then((res) => {
        console.log("typeof(res)", typeof(res))
        if(typeof(res) == "object") {
          return res.json()
        } else  {
          return res
        }
      })
      .then((data) => {
        if(data === "No results obtained"){
          setArticles([])
        }
        setArticles(data)
      })
    }
  }
  async function getBookmarks() {
   const data:any = await fetch(`http://localhost:9000/bookmarks`, {method: 'get', mode: 'cors'}).then((res) => res.json())
   
   const result = data.map((item :any) => item.Url)
   setbookmarks(result)
   
  }

  useEffect(()=>{
    fetch(`http://localhost:9000?q=`, {method: 'get', mode: 'cors'})
      .then((res) => res.json())
      .then((data) => setArticles(data))
      
  }, [])

 

  useEffect(()=> {
    getBookmarks()
  },[])

  return (
    <div className="searchparent">
      <header className='searchheader'>
        <div className='searchbox'>
            <SearchSvg className="searchsvg"/>
            <input type="text" placeholder="Search links" onChange={handleChange} value={searchText} onKeyUp={searchLinks}/>
            <button className={clearClass} type="button" title="Clear query" onClick={clearClick}>
            <Clear />
            </button>
        </div>
    </header>
    <div className='cards'>   
        {
          articles.map((card: any, index: number) => <Card key={`index-${index}`} card={card} bookmarkSet={bookmarks.includes(card.website)} />)
        }
    </div>
    </div>
  )
}

export default Search