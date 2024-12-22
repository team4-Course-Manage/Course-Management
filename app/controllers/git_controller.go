package controllers

import (
	"Course-Management/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GitController struct {
	GitService *services.GitService
}

func NewGitController(gitService *services.GitService) *GitController {
	return &GitController{
		GitService: gitService,
	}
}

func (c *GitController) CreateRepository(ctx *gin.Context) {
	var request struct {
		RepoName string `json:"repo_name" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	repoURL, err := c.GitService.CreateRepository(request.RepoName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"repo_url": repoURL})
}

func (c *GitController) GetRepository(ctx *gin.Context) {
	repoName := ctx.Param("repo_name")

	repo, err := c.GitService.GetRepository(repoName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"repository": repo})
}

func (c *GitController) DeleteRepository(ctx *gin.Context) {
	repoName := ctx.Param("repo_name")

	if err := c.GitService.DeleteRepository(repoName); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "repository deleted"})
}

func (c *GitController) ListRepositories(ctx *gin.Context) {
	repos, err := c.GitService.ListRepositories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"repositories": repos})
}

func (c *GitController) AddCollaborator(ctx *gin.Context) {
	var request struct {
		Collaborator string `json:"collaborator" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	repoName := ctx.Param("repo_name")

	if err := c.GitService.AddCollaborator(repoName, request.Collaborator); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "collaborator added"})
}

func (c *GitController) ListCommits(ctx *gin.Context) {
	repoName := ctx.Param("repo_name")

	commits, err := c.GitService.ListCommits(repoName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"commits": commits})
}
