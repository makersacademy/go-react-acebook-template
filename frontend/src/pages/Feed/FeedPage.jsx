import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";

import { getPosts } from "../../services/posts";
import { createPosts } from "../../services/posts";
import Post from "../../components/Post/Post";

import "./FeedPage.scss";

export const FeedPage = () => {
  const [posts, setPosts] = useState([]);
  const [post, setPost] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      getPosts(token)
        .then((data) => {
          setPosts(data.posts);
          localStorage.setItem("token", data.token);
        })
        .catch((err) => {
          console.error(err);
          navigate("/login");
        });
    }
  }, [navigate]);

  const token = localStorage.getItem("token");
  if (!token) {
    navigate("/login");
    return;
  }

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      await createPosts(token, post);
      const updatedPosts = await getPosts(token);
      setPosts(updatedPosts.posts);
      setPost("");
      localStorage.setItem("token", updatedPosts.token);
    } catch (err) {
      console.error(err);
    }
  }

  const handlePostChange = (event) => {
    setPost(event.target.value);
  }

  return (
    <>
    <div className="container">

      <h2>Posts</h2>
      <div className="feed" role="feed">
        {posts.map((post) => (
          <Post post={post} key={post._id} />
        ))}
      </div>

      <form onSubmit={handleSubmit}>
        <div className="create-post">
          <input
            type="text" 
            value={post} 
            onChange={handlePostChange}/>
          <input role="submit-button" id="submit" type="submit" value="Submit" />
        </div>
      </form>
        
    </div>
      
    </>
  );
};
