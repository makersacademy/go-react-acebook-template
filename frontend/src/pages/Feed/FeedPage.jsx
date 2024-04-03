import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { getPosts } from "../../services/posts";
// import { getComments } from "../../services/comments";
import Post from "../../components/Post/Post";
import LikeButton from "../../components/LikeButton/LikeButton";
import CommentsBox from "../../components/CommentsBox/CommentsBox";
import { getLikes, likeCreate, unlikeCreate } from "../../services/likes";
import "/src/FeedPage.css";

export const FeedPage = () => {
  const [posts, setPosts] = useState([]);
  const [likes, setLikes] = useState([]);
  const [liked, setLiked] = useState([]);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  // New addition - set Comments
  // const [comments, setComments] = useState([]);

  const token = localStorage.getItem("token");
  if (!token) {
    navigate("/login");
    return;
  }

  useEffect(() => {
    const fetchData = async () => {
      try {
        const token = localStorage.getItem("token");
        // Fetch posts
        const { posts: postsData } = await getPosts(token);
        setPosts(postsData);

        // Fetch likes for each post
        const likesData = {};
        const likedData = {};
        for (const post of postsData) {
          const likeData = await getLikes(post._id, token);
          likesData[post._id] = likeData.LikeCount;
          likedData[post._id] = likeData.UserHasLiked;
        }

        setLikes(likesData);
        setLiked(likedData);
      } catch (error) {
        console.error(error);
        navigate("/login");
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [navigate]);

  if (loading) {
    return <div>Loading...</div>; // Render loading indicator
  }

  const toggleLike = async (post_id) => {
    try {
      const token = localStorage.getItem("token");
      if (!token) {
        throw new Error("Token not found");
      }

      const isLiked = liked[post_id];
      if (isLiked) {
        await unlikeCreate(post_id, token);
        setLikes((prevLikes) => ({
          ...prevLikes,
          [post_id]: prevLikes[post_id] - 1,
        }));
      } else {
        await likeCreate(post_id, token);
        setLikes((prevLikes) => ({
          ...prevLikes,
          [post_id]: prevLikes[post_id] + 1,
        }));
      }
      setLiked((prevLiked) => ({
        ...prevLiked,
        [post_id]: !prevLiked[post_id],
      }));
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <>
      <h1>Posts</h1>
      <div className="feed" role="feed">
        {posts.filter(post => !post.deletedAt).map((post) => (
          <div className="post-object" key={post._id}>
            <div className="post-css">
              <Post post={post} />
            </div>
            <div className="like-css">
              <LikeButton
                postid={post._id}
                liked={liked[post._id]}
                likes={likes[post._id]}
                onToggleLike={() => toggleLike(post._id)}
              />
            </div>
            <div className="commentsbox-css">
              {/* Attempting to insert comments input box, and display comments associated with a Post ID */}
              <CommentsBox key={`comment-${post._id}`} />
              {/* {comments
              .filter((comment) => comment.post_id === post._id)
              .map((filteredComment) => (
                <Comment
                  comment={filteredComment}
                  key={filteredComment.comment._id}
                />
              ))} */}
            </div>
          </div>
        ))}
      </div>
    </>
  );
};
