Kebutuhan Entity DB
    - Users
        * Id : int
        * name : varchar
        * occupation : varchar
        * email : varchar
        * password_hash : varchar
        * avatar_file_name : varchar
        * role : varchar
        * token : varchar
        * created_at : datetime
        * updated_at : datetime
    
    - Campaigns
        * Id : int
        * user_id : int
        * name : varchar
        * short_description : varchar
        * description : text
        * goal_amount : int
        * current_amount : int
        * perks : text
        * backer_count : text
        * slug : varchar
        * created_at : datetime
        * updated_at : datetime

    - Campaign Images
        * id : int
        * campaign_id : int 
        * file_name : varchar
        * is_primary : boolean (tinyint)
        * created_at : datetime
        * updated_at : datetime
        
    - Transactions
        * id : int
        * campaign_id : int
        * user_id : int
        * amount : int
        * status : varchar
        * code : varchar
        * created_at : datetime
        * updated_at : datetime



Step Logic 

input -> Handler Mapping Input ke Struct -> Serivce Mapping ke Struct User (Memanggil Business Process) -> Repository save struct User ke Db (Pemanggilan Database) -> Database
 
Gambaran Proses
User input lalu , handler menangkap data user


INPUT
HANDLER : Mapping input dari user -> input struct
SERVICE : Melakukan mapping dari struct input ke struct user
REPOSITORY
DB

//git tutor
cara memasukan ke git 
1. Git Init terlebih dahulu
2. git add . 
3. git -m "commit nya"
4. git remote add origin https://github.com/nocturnalcoders/BackendEkost.git