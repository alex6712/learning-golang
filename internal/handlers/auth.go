package handlers

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
    var input struct { Email, Password string }
    if err := c.BindJSON(&input); err != nil {
        c.JSON(400, gin.H{"error": "Invalid input"})
        return
    }

    user, err := repository.GetUserByEmail(input.Email)
    if err != nil || !CheckPassword(input.Password, user.Password) {
        c.JSON(401, gin.H{"error": "Invalid credentials"})
        return
    }

    token, _ := utils.GenerateToken(user.ID, config.JWT.Secret)
    c.JSON(200, gin.H{"token": token})
}
