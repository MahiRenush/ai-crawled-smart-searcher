import React, {useEffect,useState} from "react";

function Bookmarks() {

    const [bookmark,setBookmark] = useState([]);
    const [title,setTitle] = useState('');
    const [url,setUrl] = useState('');
    const [loading,setLoading] = useState(false);

    useEffect(()=>{
        if(!loading){
            fetch(`http://localhost:9000/bookmarks`, {method: 'get', mode: 'cors'})
                .then((res) => res.json())
                .then((data) => setBookmark(data))
        }
      }, [loading])

      function submithandeler(event:any){
        event.preventDefault()
        setLoading(true)
        let data = {
            Title : title,
            Url : url
        }
        const requestOptions:any = {
            
            mode : 'no-cors',
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(data),
        };
        fetch('http://localhost:9000/create',requestOptions)
            .then((res) => {
                setLoading(false)
                setTitle('')
                setUrl('')
            })
      }

console.log(bookmark)
  return (
    <div className="bookmarks">
        <h3>Bookmarks</h3>

        <div className="add-bookmarks">
            <div>
                <h4 style={{letterSpacing : "4px", textAlign : "center"}}>Add Bookmark</h4>
                <form onSubmit={submithandeler}>
                    <div>
                        <label htmlFor="title">Title :</label><br/>
                        <input type="text" id="title" required placeholder="Title" name="title" value={title} onChange = {(e) => setTitle(e.target.value)}/>
                    </div>
                    <div>
                        <label htmlFor="url">URL :</label><br/>
                        <input type="url" id="url" required placeholder="URL" name="url" value={url} onChange = {(e)=> setUrl(e.target.value)} />
                    </div>
                    <button disabled = {loading} type="submit">Add</button>
            </form>
            </div>
        </div>

        <div className="bookmarks-ctn">
            <div className="bookmarks-data">
                <h5><b>Title</b></h5>
                {
                    bookmark?.map((data:any,index) => <h6 key={index}>{data.Title}</h6>)
                }
                
            </div>
            <div className="bookmarks-data">
                <h5><b>Url</b></h5>
                {
                    bookmark?.map((data:any,index) => 
                     <h6 key={index}>
                        <a href={data.Url} target = "_blank" rel="noreferrer">{data.Url}</a>
                     </h6>)
                }
            </div>
        </div> 
    </div>
  )
}

export default Bookmarks