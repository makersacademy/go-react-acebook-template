// import React, {useEffect} from 'react';
// import {getComments} from "../../services/comments.js";
import "./Post.css"
import image from "/src/static/img/x-button.png";
import userImage from "/src/static/img/user_image.png";

const Post = ({ post, onLike, user, onDelete }) => {
    const handleLikeClick = () => {
        onLike(post._id); // Passes the post ID to the parent component's like handler
    };
    const handleDeleteClick = () => {
        onDelete(post._id);
    };

    return (
        <div className="post-container" key={post._id}>
            <div className="post-info">
                <div className="post-user">
                    <img className="user-image" src={userImage} alt="image" />
                    <p>{user}</p>
                </div>
                    
                <p>{post.message}</p>
                <p>Likes: {post.likes}</p>
                <button onClick={handleLikeClick}>Like</button> {/* Like button */}
            </div>
            <div className="post-image">
                <img className="delete-button" src={image} onClick={handleDeleteClick} />
            </div>
        </div>
    );
};

export default Post;