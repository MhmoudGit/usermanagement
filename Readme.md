# WorkSpace Application

- Users get to sign up and be admin of their own workspace
- Admin can create new users in his workspace
- New users that are not admins can login without registering

### ToDo

- [x] Login Route `"/login"` Get - Post
- [x] Logout Route `"/logout"` Get
- [x] Home Route `"/"` Get

  - Home router is dynamic based on login state
  - Save user data in cookies

- [x] Signup Route `"/signup"` Get - Post

  - Workspace automatically created for this user
  - Registered user is admin of that workspace

- [ ] Workspace Management
- [ ] User CRUD
