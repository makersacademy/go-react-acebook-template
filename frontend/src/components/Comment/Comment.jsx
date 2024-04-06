

// import image from "/src/static/img/x-button.png";
import userImage from "/src/static/img/user_image.png";

const Comment= ({comment, onDelete}) => {

    const handleDeleteCommentClick = () => {
        onDelete(comment._id);
    };

    return (
        <div>
                <div className="comment-info">
                    <div className="comment-user">
                        <img className="user-image" src={userImage} alt="image" />
                        {/* <p>{post.User.username}</p> */}
                    </div>
                    <p>{comment.message}</p>
                    <button onClick={handleDeleteCommentClick}>Delete</button>
                </div>
        </div>
    );
};

export default Comment;