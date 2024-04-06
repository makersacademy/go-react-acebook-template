import "./Post.css";
// import image from "/src/static/img/x-button.png";
import { useState, useEffect} from "react";
import Comment from "../Comment/Comment";
import {createComments, getComments, deleteComments} from "../../services/comments"

const Post = ({ post, onLike, user, onDelete, token}) => {
const [comments, setComments] = useState([]);
const [comment, setComment] = useState("");


const handleLikeClick = () => {
onLike(post._id);
};

const handleDeleteClick = () => {
onDelete(post._id);
};

useEffect(() => {
const token = localStorage.getItem("token");
if (token) {
    getComments(post._id, token)
        .then((data) => {
            console.log(data)
            const sortedPosts = data.comments.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
            setComments(sortedPosts)
            localStorage.setItem("token", data.token);
        })
        .catch((err) => {
            console.error(err);
        });
}
}, [post._id]);


const handleSubmitComment = async (event) => {
event.preventDefault();
    try {
        console.log("Token:", token);
        console.log("Comment:", comment);
        console.log("Post ID:", post._id);
        await createComments(token, comment, post._id);
        const updatedComments = await getComments(post._id, token);
        const sortedComments = updatedComments.comments.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
        setComments(sortedComments);
        setComment("");
        localStorage.setItem("token", updatedComments.token);
    } catch (err) {
        console.error(err);
    }
};

const handleDeleteComment = async (commentId) => {
try {
    await deleteComments(token, post._id, commentId);
    const updatedComments = await getComments(post._id, token);
        const sortedComments = updatedComments.comments.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
        setComments(sortedComments);
} catch (err) {
    console.error(err);
}
};

const handleCommentChange = (event) => {
setComment(event.target.value);
};

return (
<div className="post-container" key={post._id}>
    <div className="post-info">
        <div className="post-user">
            <img className="user-image" src={user.image} alt="image" />
            <p>{user.username}</p>
        </div>
    <div>

    </div>
    <p>{post.message}</p>
    <p>Likes: {post.likes}</p>
    <button onClick={handleLikeClick}>Like</button>
    <button onClick={handleDeleteClick}>Delete</button>
    </div>

    <form onSubmit={handleSubmitComment}>
    <div className="create-comment">
        <input
        type="text"
        onChange={handleCommentChange}
        placeholder="Add a comment..."
        />
        <button type="submit">Submit</button>
    </div>
    </form>

    {comments
    .map((comment) => (
        <div className="feed-comment" key={comment._id}>
        <Comment comment={comment} onDelete={handleDeleteComment}/>
        </div>
    ))}
</div>
);
};

export default Post;
