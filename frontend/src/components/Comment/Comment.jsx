import { useState } from "react";

const Comment = (props) => {
  const [comment, setComment] = useState("");
  const [comments, setComments] = useState(props.post.comments || []);

  const handleCommentChange = (event) => {
    setComment(event.target.value);
  };

  const handleSubmitComment = (event) => {
    event.preventDefault();
    // You can handle submitting the comment here
    console.log("Submitting comment:", comment);
    const newComment = { text: comment, id: Math.random().toString() }; // You can replace id generation with a better approach
    setComments([...comments, newComment]); // Add the new comment to the list of comments
    setComment(""); // Clear the comment input field after submission
  };

  return (
    <div>
      <h4>Comments:</h4>
      <ul>
        {comments.map((comment) => (
          <li key={comment.id}>{comment.text}</li>
        ))}
      </ul>
      <form onSubmit={handleSubmitComment}>
        <input
          type="text"
          value={comment}
          onChange={handleCommentChange}
          placeholder="Add a comment..."
        />
        <button type="submit">Submit</button>
      </form>
    </div>
  );
};

export default Comment;
