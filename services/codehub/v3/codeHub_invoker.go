package v3

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/services/codehub/v3/model"
)

type CreateCommitInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCommitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCommitInvoker) Invoke() (*model.CreateCommitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCommitResponse), nil
	}
}

type ListCommitsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCommitsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCommitsInvoker) Invoke() (*model.ListCommitsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCommitsResponse), nil
	}
}

type ShowDiffCommitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDiffCommitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDiffCommitInvoker) Invoke() (*model.ShowDiffCommitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDiffCommitResponse), nil
	}
}

type ShowSingleCommitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSingleCommitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSingleCommitInvoker) Invoke() (*model.ShowSingleCommitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSingleCommitResponse), nil
	}
}

type CreateMergeRequestDiscussionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMergeRequestDiscussionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMergeRequestDiscussionInvoker) Invoke() (*model.CreateMergeRequestDiscussionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMergeRequestDiscussionResponse), nil
	}
}

type CreateMergeRequestDiscussionNoteInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMergeRequestDiscussionNoteInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMergeRequestDiscussionNoteInvoker) Invoke() (*model.CreateMergeRequestDiscussionNoteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMergeRequestDiscussionNoteResponse), nil
	}
}

type ShowReviewSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowReviewSettingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowReviewSettingInvoker) Invoke() (*model.ShowReviewSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowReviewSettingResponse), nil
	}
}

type ListFilesByQueryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFilesByQueryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFilesByQueryInvoker) Invoke() (*model.ListFilesByQueryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFilesByQueryResponse), nil
	}
}

type ShowFileInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowFileInvoker) Invoke() (*model.ShowFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFileResponse), nil
	}
}

type GetAllRepositoryByProjectIdInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *GetAllRepositoryByProjectIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *GetAllRepositoryByProjectIdInvoker) Invoke() (*model.GetAllRepositoryByProjectIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetAllRepositoryByProjectIdResponse), nil
	}
}

type GetProductTemplatesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *GetProductTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *GetProductTemplatesInvoker) Invoke() (*model.GetProductTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetProductTemplatesResponse), nil
	}
}

type ListProductTwoTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProductTwoTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProductTwoTemplatesInvoker) Invoke() (*model.ListProductTwoTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProductTwoTemplatesResponse), nil
	}
}

type ShowRepositoryNameExistInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepositoryNameExistInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepositoryNameExistInvoker) Invoke() (*model.ShowRepositoryNameExistResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepositoryNameExistResponse), nil
	}
}

type AddRepoMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddRepoMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddRepoMembersInvoker) Invoke() (*model.AddRepoMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddRepoMembersResponse), nil
	}
}

type DeleteRepoMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRepoMemberInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRepoMemberInvoker) Invoke() (*model.DeleteRepoMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRepoMemberResponse), nil
	}
}

type ListRepoMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepoMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepoMembersInvoker) Invoke() (*model.ListRepoMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepoMembersResponse), nil
	}
}

type SetRepoRoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetRepoRoleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetRepoRoleInvoker) Invoke() (*model.SetRepoRoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetRepoRoleResponse), nil
	}
}

type AddDeployKeyInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *AddDeployKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *AddDeployKeyInvoker) Invoke() (*model.AddDeployKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddDeployKeyResponse), nil
	}
}

type AddDeployKeyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *AddDeployKeyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddDeployKeyV2Invoker) Invoke() (*model.AddDeployKeyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddDeployKeyV2Response), nil
	}
}

type AddProtectBranchV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *AddProtectBranchV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddProtectBranchV2Invoker) Invoke() (*model.AddProtectBranchV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddProtectBranchV2Response), nil
	}
}

type AddTagV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *AddTagV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddTagV2Invoker) Invoke() (*model.AddTagV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddTagV2Response), nil
	}
}

type CreateRepositoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRepositoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRepositoryInvoker) Invoke() (*model.CreateRepositoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRepositoryResponse), nil
	}
}

type DeleteDeployKeyInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteDeployKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteDeployKeyInvoker) Invoke() (*model.DeleteDeployKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDeployKeyResponse), nil
	}
}

type DeleteDeployKeyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDeployKeyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDeployKeyV2Invoker) Invoke() (*model.DeleteDeployKeyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDeployKeyV2Response), nil
	}
}

type DeleteRepositoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRepositoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRepositoryInvoker) Invoke() (*model.DeleteRepositoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRepositoryResponse), nil
	}
}

type GetRepositoryByProjectIdInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *GetRepositoryByProjectIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *GetRepositoryByProjectIdInvoker) Invoke() (*model.GetRepositoryByProjectIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetRepositoryByProjectIdResponse), nil
	}
}

type GetTemplatesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *GetTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *GetTemplatesInvoker) Invoke() (*model.GetTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetTemplatesResponse), nil
	}
}

type ListBranchesByRepositoryIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBranchesByRepositoryIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBranchesByRepositoryIdInvoker) Invoke() (*model.ListBranchesByRepositoryIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBranchesByRepositoryIdResponse), nil
	}
}

type ListCommitStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCommitStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCommitStatisticsInvoker) Invoke() (*model.ListCommitStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCommitStatisticsResponse), nil
	}
}

type ListFilesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFilesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFilesInvoker) Invoke() (*model.ListFilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFilesResponse), nil
	}
}

type ListMergeChangesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMergeChangesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMergeChangesInvoker) Invoke() (*model.ListMergeChangesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMergeChangesResponse), nil
	}
}

type ListMergeChangesTreesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMergeChangesTreesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMergeChangesTreesInvoker) Invoke() (*model.ListMergeChangesTreesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMergeChangesTreesResponse), nil
	}
}

type ListMergeRequestInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMergeRequestInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMergeRequestInvoker) Invoke() (*model.ListMergeRequestResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMergeRequestResponse), nil
	}
}

type ListMergeRequestReviewersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMergeRequestReviewersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMergeRequestReviewersInvoker) Invoke() (*model.ListMergeRequestReviewersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMergeRequestReviewersResponse), nil
	}
}

type ListRelatedCommitsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRelatedCommitsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRelatedCommitsInvoker) Invoke() (*model.ListRelatedCommitsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRelatedCommitsResponse), nil
	}
}

type ListRepositoryStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepositoryStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepositoryStatusInvoker) Invoke() (*model.ListRepositoryStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepositoryStatusResponse), nil
	}
}

type ListSubfilesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubfilesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubfilesInvoker) Invoke() (*model.ListSubfilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubfilesResponse), nil
	}
}

type ListTemplatesTwoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTemplatesTwoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTemplatesTwoInvoker) Invoke() (*model.ListTemplatesTwoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTemplatesTwoResponse), nil
	}
}

type ListTwoTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTwoTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTwoTemplatesInvoker) Invoke() (*model.ListTwoTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTwoTemplatesResponse), nil
	}
}

type ShareTemplatesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShareTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShareTemplatesInvoker) Invoke() (*model.ShareTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShareTemplatesResponse), nil
	}
}

type ShowBranchesByRepositoryIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBranchesByRepositoryIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBranchesByRepositoryIdInvoker) Invoke() (*model.ShowBranchesByRepositoryIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBranchesByRepositoryIdResponse), nil
	}
}

type ShowBranchesByTwoRepositoryIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBranchesByTwoRepositoryIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBranchesByTwoRepositoryIdInvoker) Invoke() (*model.ShowBranchesByTwoRepositoryIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBranchesByTwoRepositoryIdResponse), nil
	}
}

type ShowCommitsByBranchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCommitsByBranchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCommitsByBranchInvoker) Invoke() (*model.ShowCommitsByBranchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCommitsByBranchResponse), nil
	}
}

type ShowCommitsByRepoIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCommitsByRepoIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCommitsByRepoIdInvoker) Invoke() (*model.ShowCommitsByRepoIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCommitsByRepoIdResponse), nil
	}
}

type ShowHasPipelineInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHasPipelineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHasPipelineInvoker) Invoke() (*model.ShowHasPipelineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHasPipelineResponse), nil
	}
}

type ShowImageBlobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageBlobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowImageBlobInvoker) Invoke() (*model.ShowImageBlobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageBlobResponse), nil
	}
}

type ShowMasterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMasterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMasterInvoker) Invoke() (*model.ShowMasterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMasterResponse), nil
	}
}

type ShowMergeRequestInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMergeRequestInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMergeRequestInvoker) Invoke() (*model.ShowMergeRequestResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMergeRequestResponse), nil
	}
}

type ShowRepoIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepoIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepoIdInvoker) Invoke() (*model.ShowRepoIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepoIdResponse), nil
	}
}

type ShowRepositoryArchiveInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepositoryArchiveInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepositoryArchiveInvoker) Invoke() (*model.ShowRepositoryArchiveResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepositoryArchiveResponse), nil
	}
}

type ShowRepositoryByUuidInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepositoryByUuidInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepositoryByUuidInvoker) Invoke() (*model.ShowRepositoryByUuidResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepositoryByUuidResponse), nil
	}
}

type ShowRepositoryStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRepositoryStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRepositoryStatisticsInvoker) Invoke() (*model.ShowRepositoryStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRepositoryStatisticsResponse), nil
	}
}

type ShowStatisticCommitInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowStatisticCommitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowStatisticCommitInvoker) Invoke() (*model.ShowStatisticCommitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStatisticCommitResponse), nil
	}
}

type ShowStatisticCommitV3Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStatisticCommitV3Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStatisticCommitV3Invoker) Invoke() (*model.ShowStatisticCommitV3Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStatisticCommitV3Response), nil
	}
}

type ShowStatisticalDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStatisticalDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStatisticalDataInvoker) Invoke() (*model.ShowStatisticalDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStatisticalDataResponse), nil
	}
}

type UpdateMergeRequestApprovalStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMergeRequestApprovalStateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMergeRequestApprovalStateInvoker) Invoke() (*model.UpdateMergeRequestApprovalStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMergeRequestApprovalStateResponse), nil
	}
}

type AddSshKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddSshKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddSshKeyInvoker) Invoke() (*model.AddSshKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddSshKeyResponse), nil
	}
}

type DeleteSShkeyInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteSShkeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteSShkeyInvoker) Invoke() (*model.DeleteSShkeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSShkeyResponse), nil
	}
}

type ListSshKeysInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSshKeysInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSshKeysInvoker) Invoke() (*model.ListSshKeysResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSshKeysResponse), nil
	}
}

type ShowPrivateKeyVerifyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPrivateKeyVerifyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPrivateKeyVerifyInvoker) Invoke() (*model.ShowPrivateKeyVerifyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPrivateKeyVerifyResponse), nil
	}
}

type ValidateHttpsInfoInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ValidateHttpsInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ValidateHttpsInfoInvoker) Invoke() (*model.ValidateHttpsInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateHttpsInfoResponse), nil
	}
}

type ValidateHttpsInfoV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateHttpsInfoV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ValidateHttpsInfoV2Invoker) Invoke() (*model.ValidateHttpsInfoV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateHttpsInfoV2Response), nil
	}
}

type CreateNewBranchInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNewBranchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateNewBranchInvoker) Invoke() (*model.CreateNewBranchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNewBranchResponse), nil
	}
}

type AssociateIssuesInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateIssuesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateIssuesInvoker) Invoke() (*model.AssociateIssuesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateIssuesResponse), nil
	}
}

type CreateProjectAndRepositoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProjectAndRepositoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateProjectAndRepositoriesInvoker) Invoke() (*model.CreateProjectAndRepositoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProjectAndRepositoriesResponse), nil
	}
}

type CreateProjectAndforkRepositoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProjectAndforkRepositoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateProjectAndforkRepositoriesInvoker) Invoke() (*model.CreateProjectAndforkRepositoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProjectAndforkRepositoriesResponse), nil
	}
}

type ListUserAllRepositoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserAllRepositoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUserAllRepositoriesInvoker) Invoke() (*model.ListUserAllRepositoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserAllRepositoriesResponse), nil
	}
}

type ShowAllRepositoryByTwoProjectIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAllRepositoryByTwoProjectIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAllRepositoryByTwoProjectIdInvoker) Invoke() (*model.ShowAllRepositoryByTwoProjectIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAllRepositoryByTwoProjectIdResponse), nil
	}
}

type AddHooksInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddHooksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddHooksInvoker) Invoke() (*model.AddHooksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddHooksResponse), nil
	}
}

type DeleteHooksInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHooksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHooksInvoker) Invoke() (*model.DeleteHooksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHooksResponse), nil
	}
}

type ListHooksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHooksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHooksInvoker) Invoke() (*model.ListHooksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHooksResponse), nil
	}
}
