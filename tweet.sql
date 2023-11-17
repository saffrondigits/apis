CREATE TABLE tweets (
  id serial PRIMARY KEY, -- a unique identifier for each tweet
  user_id integer NOT NULL, -- a reference to the user who posted the tweet
  content text NOT NULL, -- the text content of the tweet
  created_at timestamp NOT NULL, -- the date and time when the tweet was posted
  likes integer DEFAULT 0, -- the number of likes the tweet received
  retweets integer DEFAULT 0, -- the number of retweets the tweet received
  FOREIGN KEY (user_id) REFERENCES users (id) -- a foreign key constraint to link the tweet to the user
);