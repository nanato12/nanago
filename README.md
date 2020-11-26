# nanago
[Nana](https://nana-music.com/) Unofficial API

## go get
```bash
$ go get github.com/nanato12/nanago
```

## import
```go
import (
    "github.com/nanato12/nanago/nana"
)
```

## CreateAccount
Create new Nana account.  
Email address and password will be randomly determined when empty.

```go
client, err := nana.CreateAccount(mailAdress, password)
```

## Login
There are two ways to log in.
- MailAddress & Password Login.
```go
client, err := nana.Login(mailAddress, password)
```
- Token Login
```go
client, err := nana.LoginByToken(token)
```

## Client Information
```go
fmt.Println("[ID]", client.ID, "[NAME]", client.Name, "[TOKEN]", client.Token)
```

## Account Information
```go
res, err := client.GetMyInfo()
fmt.Println(res)
```

## Follow User 🙆
Follow user by `UserID`.
```go
res, err := client.Follow(userID)
```

## Play Post ▶️
Play post by `PostID`
```go
res, err := client.PlayPost(postID)
```
## Applause Post 👏
Applause post by `PostID`
```go
res, err := client.ApplausePost(postID)
```
## Comment on Post 💬
Comment on post by `PostID`
```go
res, err := client.CommentPost(postID, comment)
```
## Join Community 👨‍👨‍👧‍👧
Join community by `CommunityID`
```go
res, err := client.JoinCommunity(communityID)
```
## Sample
Create account and follow user.
```go
const (
	YourUserID = "XXXXXXXXXX"
)

func main() {
	client, err := nana.CreateAccount("", "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("[ID]", client.ID, "[NAME]", client.Name)
	if _, err := client.Follow(YourUserID); err != nil {
		fmt.Println(err)
		return
	}
}
```
