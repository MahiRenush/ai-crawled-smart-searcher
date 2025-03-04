import {useState, useEffect} from "react"
import Card from "./Card";

function MostUpvoted() {
  const [articles, setArticles] =  useState<any>([]);

  useEffect(()=>{
    fetch(`http://localhost:9000?q=`, {method: 'get', mode: 'cors'})
      .then((res) => res.json())
      .then((data) => setArticles(data))
  }, [])
  
  return (
    <div className="popular">
      <h3>Most Upvoted</h3>
      <div className="cards">
        {
          articles.sort((a: any, b:any ) => b.upvotes - a.upvotes)
          .map((card: any, index: number) => <Card key={`index-${index}`}card={card} />)
        }
      </div>
    </div>
  );
}

export default MostUpvoted;
