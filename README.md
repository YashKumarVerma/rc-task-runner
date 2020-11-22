# Reverse Coding : Task Runner

![Build package](https://github.com/YashKumarVerma/rc-task-runner/workflows/Build%20package/badge.svg)

The service responsible running code, feeding input, and returning responses.
- API Docs: [@Postman](https://documenter.getpostman.com/view/10043948/TVev4k25) 

![https://i.imgur.com/ubXnojN.png](https://i.imgur.com/ubXnojN.png)

## Known Issues
- Response too large: depending on client and connection, some responses might be too large to be returned to the user and might cause issues. As of now, the system does not return output on them, and following the principle of reliability, **continues to send responses to other clients connected**.

![https://i.imgur.com/9HUhdOr.png](https://i.imgur.com/9HUhdOr.png)