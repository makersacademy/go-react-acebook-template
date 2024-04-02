import React, {useEffect} from 'react';
import {getComments} from "../../services/comments.js";

const Post = ({ post, onLike }) => {
    const handleLikeClick = () => {
        onLike(post._id); // Passes the post ID to the parent component's like handler
    };

    return (
        <article key={post._id}>
            <p>{post.message}</p>
            <p>Likes: {post.likes}</p>
            <button onClick={handleLikeClick}>Like</button> {/* Like button */}
        </article>
    );
};

export default Post;