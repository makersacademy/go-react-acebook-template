import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { getPosts, createPosts, updatePostLikes} from "../../services/posts";
import Post from "../../components/Post/Post";
import Comment from "../../components/Comment/Comment";
import { getComments, createComments } from "../../services/comments";
import "./FeedPage.scss"



export const FeedPage = () => {
    const [posts, setPosts] = useState([]);
    const [post, setPost] = useState("");
    const [comments, setComments] = useState([]);
    const navigate = useNavigate();

    useEffect(() => {
        const token = localStorage.getItem("token");
        if (token) {
            getPosts(token)
                .then((data) => {
                    const sortedPosts = data.posts.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
                    setPosts(sortedPosts);
                    localStorage.setItem("token", data.token);
                })
                .catch((err) => {
                    console.error(err);
                    navigate("/login");
                });
            getComments(token)
                .then((data) => {
                    const sortedComments = data.comments.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
                    setComments(sortedComments);
                    localStorage.setItem("token", data.token);
                })
                .catch((err) => {
                    console.error(err);
                });
        }
    }, [navigate]);

    const handleSubmit = async (event) => {
        event.preventDefault();
        try {
            await createPosts(token, post);
            const updatedPosts = await getPosts(token);
            const sortedPosts = updatedPosts.posts.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
            setPosts(sortedPosts);
            setPost("");
            localStorage.setItem("token", updatedPosts.token);
        } catch (err) {
            console.error(err);
        }
    };

    const handleLike = async (postId) => {
        try {
            await updatePostLikes(token, postId);
            const updatedPosts = await getPosts(token);
            const sortedPosts = updatedPosts.posts.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
            setPosts(sortedPosts);
        } catch (err) {
            console.error(err);
        }
    };
  const token = localStorage.getItem("token");
  if (!token) {
    navigate("/login");
    return;
  }

  const handleSubmitPost = async (event) => {
    event.preventDefault();
    try {
      const createdPostResponse = await createPosts(token, post);
      const updatedPosts = await getPosts(token);
      const sortedPosts = updatedPosts.posts.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
      setPosts(sortedPosts);
      setPost("");
      localStorage.setItem("token", updatedPosts.token);
    } catch (err) {
      console.error(err);
    }
  };

  const handleSubmitComment = async (postId, comment) => {
    try {
      const CommentResponse = await createComments(token, postId, comment);
      const updatedComments = await getComments(token);
      const sortedComments = updatedComments.comments.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
      setComments(sortedComments);
      localStorage.setItem("token", CommentResponse.token);
    } catch (err) {
      console.error(err);
    }
  };

  const handlePostChange = (event) => {
    setPost(event.target.value);
  };

  return (
    <div className="feed-container">
      <h2>Posts</h2>
      <div className="feed-all-posts" role="feed">
        {posts.map((post) => (
          <div className="feed-post" key={post._id}>
            <Post post={post} onLike={handleLike} />
            <Comment post={post} comments={comments.filter((comment) => comment.postId === post._id)} onSubmit={(comment) => handleSubmitComment(post._id, comment)} />
          </div>
        ))}
      </div>
      <form onSubmit={handleSubmitPost}>
        <div className="create-post">
          <input type="text" value={post} onChange={handlePostChange} />
          <input role="submit-button" id="submit" type="submit" value="Submit" />
        </div>
      </form>
    </div>
  )
};
