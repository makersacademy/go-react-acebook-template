// docs: https://vitejs.dev/guide/env-and-mode.html
const BACKEND_URL = import.meta.env.VITE_BACKEND_URL;

export const getComments = async (postId, token) => {
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${token}`,
    },
  };

  const response = await fetch(`${BACKEND_URL}/posts/${postId}/comments`, requestOptions);

  if (response.status !== 200) {
    throw new Error("Unable to fetch comments");
  }

  const data = await response.json();
  return data;
};

export const createComments = async (token, comment, postId) => {
  const commentData = {
    "message": comment
  };
  
  const requestOptions = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify(commentData),
  };

  const response = await fetch(`${BACKEND_URL}/posts/${postId}/comments`, requestOptions);

  if (response.status !== 201) {
    throw new Error("Unable to create comment");
  }

  const data = await response.json();
  return data;
};


export const deleteComments = async (token, postId, commentId) => {
  const requestOptions = {
    method: "DELETE",
    headers: {
      Authorization: `Bearer ${token}`,
    },
  };

  const response = await fetch(`${BACKEND_URL}/posts/${postId}/comments/${commentId}/delete`, requestOptions);
  if (response.status !== 200) {
    throw new Error("Unable to fetch posts");
  }

  const data = await response.json();
  return data;
};