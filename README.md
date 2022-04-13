# TansSactional

So how does things work?
* Import Postman Requests [Curl](https://www.getpostman.com/collections/7b0cb4ae55f7e9fe3020)
(Import using JSON link)
## Things I used:-
* Go(lang)
* SQL
* JWT

## Explaining APIs:-
* Get Username
  - Input:Username
  - Output:userId,firstName,lastName,penName,userEmail,bio,number
  - Sample URL:- https://userserviceshivam.herokuapp.com/api/v0/getUser/:username
* Create User
  - Input:userId,firstName,lastName,penName,userEmail,bio,number,password
  - Output:token
  - Sample URL:- https://userserviceshivam.herokuapp.com/api/v0/createUser
* Login User
  - Input:penName,password
  - Output:token
  - Sample URL:- https://userserviceshivam.herokuapp.com/api/v0/loginUser
* Fetch Unlocked Chapters
  - Input:UserID,SeriesID
  - Output:[]{seriesId,title,story}
  - Sample URL:-https://contentserviceshivam.herokuapp.com/api/v0/content/fetchContent/:userID/:seriesID
* Bulk Upload Chapters
  - Input:[]{seriesId,title,story}
  - Output:Success Response
  - Ssample URL:-https://contentserviceshivam.herokuapp.com/api/v0/content/uploadContent
* Get Series
  - Input:seriesId
  - Output:seriesId,author,name,chapters
  - Sample URL:-https://contentserviceshivam.herokuapp.com/api/v0/content/getSeries/:seriesID
* Create Series
  - Input:author,name,chapters
  - Output:Success Response
  - Sample URL:-https://contentserviceshivam.herokuapp.com/api/v0/content/uploadSeries
* Unlock Chapters
  - Input:UserID,SeriesID
  - Output:Success Response
  - Sample URL:-https://dailypassserviceshivam.herokuapp.com/api/v0/dailypass/unlockContent
* Get Unlock Chapters
  - Input:UserID
  - Output:Chapter and Unlocked List for given user
  - Sample URL:-https://dailypassserviceshivam.herokuapp.com/api/v0/dailypass/getUnlockedContent/:userID
 
## Overall Schematic 
![Group Schema (8)](https://user-images.githubusercontent.com/60891544/163135586-5e9e20c3-8180-4813-94c5-4ddcb2d80419.png)
## User Schematic
![System Flow (13)](https://user-images.githubusercontent.com/60891544/163108333-97ee3311-355e-4ff6-98d0-c0738263b209.png)
## Company Schematic
![Source Diagram (2)](https://user-images.githubusercontent.com/60891544/163108338-882c1501-8bce-406e-b9b3-009e45b98ab7.png)
## Invoices Schematic
![Source Diagram (3)](https://user-images.githubusercontent.com/60891544/163131509-ba6bd736-a211-4e22-9922-a5ac5e2505a6.png)
