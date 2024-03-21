package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/idmeclassicapi/v1/model"
)

type AddTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddTagInvoker) Invoke() (*model.AddTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddTagResponse), nil
	}
}

type BatchCheckinInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCheckinInvoker) Invoke() (*model.BatchCheckinResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCheckinResponse), nil
	}
}

type BatchCheckoutInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCheckoutInvoker) Invoke() (*model.BatchCheckoutResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCheckoutResponse), nil
	}
}

type BatchCheckoutAndUpdateInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCheckoutAndUpdateInvoker) Invoke() (*model.BatchCheckoutAndUpdateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCheckoutAndUpdateResponse), nil
	}
}

type BatchCheckoutUndoInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCheckoutUndoInvoker) Invoke() (*model.BatchCheckoutUndoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCheckoutUndoResponse), nil
	}
}

type BatchCheckoutUndoByAdminInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCheckoutUndoByAdminInvoker) Invoke() (*model.BatchCheckoutUndoByAdminResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCheckoutUndoByAdminResponse), nil
	}
}

type BatchCreateUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateUsingPostInvoker) Invoke() (*model.BatchCreateUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateUsingPostResponse), nil
	}
}

type BatchCreateViewInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateViewInvoker) Invoke() (*model.BatchCreateViewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateViewResponse), nil
	}
}

type BatchDeleteBranchInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteBranchInvoker) Invoke() (*model.BatchDeleteBranchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteBranchResponse), nil
	}
}

type BatchDeleteLogicalBranchInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteLogicalBranchInvoker) Invoke() (*model.BatchDeleteLogicalBranchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteLogicalBranchResponse), nil
	}
}

type BatchDeleteLogicalUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteLogicalUsingPostInvoker) Invoke() (*model.BatchDeleteLogicalUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteLogicalUsingPostResponse), nil
	}
}

type BatchDeleteUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteUsingPostInvoker) Invoke() (*model.BatchDeleteUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteUsingPostResponse), nil
	}
}

type BatchExecuteReviseInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchExecuteReviseInvoker) Invoke() (*model.BatchExecuteReviseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchExecuteReviseResponse), nil
	}
}

type BatchShowGetUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchShowGetUsingPostInvoker) Invoke() (*model.BatchShowGetUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchShowGetUsingPostResponse), nil
	}
}

type BatchUpdateAndCheckinInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateAndCheckinInvoker) Invoke() (*model.BatchUpdateAndCheckinResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateAndCheckinResponse), nil
	}
}

type BatchUpdateAndReviseInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateAndReviseInvoker) Invoke() (*model.BatchUpdateAndReviseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateAndReviseResponse), nil
	}
}

type BatchUpdateByAdminInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateByAdminInvoker) Invoke() (*model.BatchUpdateByAdminResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateByAdminResponse), nil
	}
}

type BatchUpdateUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateUsingPostInvoker) Invoke() (*model.BatchUpdateUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateUsingPostResponse), nil
	}
}

type BatchUpdateVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateVersionInvoker) Invoke() (*model.BatchUpdateVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateVersionResponse), nil
	}
}

type CheckinInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckinInvoker) Invoke() (*model.CheckinResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckinResponse), nil
	}
}

type CheckoutInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckoutInvoker) Invoke() (*model.CheckoutResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckoutResponse), nil
	}
}

type CheckoutAndUpdateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckoutAndUpdateInvoker) Invoke() (*model.CheckoutAndUpdateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckoutAndUpdateResponse), nil
	}
}

type CheckoutUndoInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckoutUndoInvoker) Invoke() (*model.CheckoutUndoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckoutUndoResponse), nil
	}
}

type CheckoutUndoByAdminInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckoutUndoByAdminInvoker) Invoke() (*model.CheckoutUndoByAdminResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckoutUndoByAdminResponse), nil
	}
}

type CollectHistoryDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *CollectHistoryDataInvoker) Invoke() (*model.CollectHistoryDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectHistoryDataResponse), nil
	}
}

type CompareBusinessVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CompareBusinessVersionInvoker) Invoke() (*model.CompareBusinessVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CompareBusinessVersionResponse), nil
	}
}

type CountUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountUsingPostInvoker) Invoke() (*model.CountUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountUsingPostResponse), nil
	}
}

type CreateMultiViewInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMultiViewInvoker) Invoke() (*model.CreateMultiViewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMultiViewResponse), nil
	}
}

type CreateUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateUsingPostInvoker) Invoke() (*model.CreateUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateUsingPostResponse), nil
	}
}

type CreateViewInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateViewInvoker) Invoke() (*model.CreateViewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateViewResponse), nil
	}
}

type DeleteBranchInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBranchInvoker) Invoke() (*model.DeleteBranchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBranchResponse), nil
	}
}

type DeleteByConditionMultiViewInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteByConditionMultiViewInvoker) Invoke() (*model.DeleteByConditionMultiViewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteByConditionMultiViewResponse), nil
	}
}

type DeleteByConditionUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteByConditionUsingPostInvoker) Invoke() (*model.DeleteByConditionUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteByConditionUsingPostResponse), nil
	}
}

type DeleteLatestVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLatestVersionInvoker) Invoke() (*model.DeleteLatestVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLatestVersionResponse), nil
	}
}

type DeleteLogicalBranchInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLogicalBranchInvoker) Invoke() (*model.DeleteLogicalBranchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLogicalBranchResponse), nil
	}
}

type DeleteLogicalLatestVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLogicalLatestVersionInvoker) Invoke() (*model.DeleteLogicalLatestVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLogicalLatestVersionResponse), nil
	}
}

type DeleteUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteUsingPostInvoker) Invoke() (*model.DeleteUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteUsingPostResponse), nil
	}
}

type DisableDataInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableDataInstanceInvoker) Invoke() (*model.DisableDataInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableDataInstanceResponse), nil
	}
}

type EnableDataInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableDataInstanceInvoker) Invoke() (*model.EnableDataInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableDataInstanceResponse), nil
	}
}

type ExecuteReviseInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteReviseInvoker) Invoke() (*model.ExecuteReviseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteReviseResponse), nil
	}
}

type GenerateBusinessCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *GenerateBusinessCodeInvoker) Invoke() (*model.GenerateBusinessCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GenerateBusinessCodeResponse), nil
	}
}

type ListAllVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllVersionsInvoker) Invoke() (*model.ListAllVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllVersionsResponse), nil
	}
}

type ListHistoryDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHistoryDataInvoker) Invoke() (*model.ListHistoryDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHistoryDataResponse), nil
	}
}

type ListQueryUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQueryUsingPostInvoker) Invoke() (*model.ListQueryUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQueryUsingPostResponse), nil
	}
}

type ListSelectUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSelectUsingPostInvoker) Invoke() (*model.ListSelectUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSelectUsingPostResponse), nil
	}
}

type ListUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUsingPostInvoker) Invoke() (*model.ListUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUsingPostResponse), nil
	}
}

type RemoveTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveTagInvoker) Invoke() (*model.RemoveTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveTagResponse), nil
	}
}

type SaveAllUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *SaveAllUsingPostInvoker) Invoke() (*model.SaveAllUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SaveAllUsingPostResponse), nil
	}
}

type SaveAsUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *SaveAsUsingPostInvoker) Invoke() (*model.SaveAsUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SaveAsUsingPostResponse), nil
	}
}

type SaveUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *SaveUsingPostInvoker) Invoke() (*model.SaveUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SaveUsingPostResponse), nil
	}
}

type ShowFindUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFindUsingPostInvoker) Invoke() (*model.ShowFindUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFindUsingPostResponse), nil
	}
}

type ShowGetByUniqueKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGetByUniqueKeyInvoker) Invoke() (*model.ShowGetByUniqueKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGetByUniqueKeyResponse), nil
	}
}

type ShowGetUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGetUsingPostInvoker) Invoke() (*model.ShowGetUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGetUsingPostResponse), nil
	}
}

type ShowLogicalDeleteByConditionUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLogicalDeleteByConditionUsingPostInvoker) Invoke() (*model.ShowLogicalDeleteByConditionUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLogicalDeleteByConditionUsingPostResponse), nil
	}
}

type ShowLogicalDeleteUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLogicalDeleteUsingPostInvoker) Invoke() (*model.ShowLogicalDeleteUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLogicalDeleteUsingPostResponse), nil
	}
}

type ShowStaticsUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStaticsUsingPostInvoker) Invoke() (*model.ShowStaticsUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStaticsUsingPostResponse), nil
	}
}

type ShowTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTagInvoker) Invoke() (*model.ShowTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTagResponse), nil
	}
}

type ShowVersionByMasterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVersionByMasterInvoker) Invoke() (*model.ShowVersionByMasterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVersionByMasterResponse), nil
	}
}

type SwitchLifecycleTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchLifecycleTemplateInvoker) Invoke() (*model.SwitchLifecycleTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchLifecycleTemplateResponse), nil
	}
}

type UpdateAndCheckinInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAndCheckinInvoker) Invoke() (*model.UpdateAndCheckinResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAndCheckinResponse), nil
	}
}

type UpdateAndReviseInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAndReviseInvoker) Invoke() (*model.UpdateAndReviseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAndReviseResponse), nil
	}
}

type UpdateByAdminInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateByAdminInvoker) Invoke() (*model.UpdateByAdminResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateByAdminResponse), nil
	}
}

type UpdateByConditionUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateByConditionUsingPostInvoker) Invoke() (*model.UpdateByConditionUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateByConditionUsingPostResponse), nil
	}
}

type UpdateStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStateInvoker) Invoke() (*model.UpdateStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStateResponse), nil
	}
}

type UpdateUsingPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUsingPostInvoker) Invoke() (*model.UpdateUsingPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUsingPostResponse), nil
	}
}

type UpdateViewInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateViewInvoker) Invoke() (*model.UpdateViewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateViewResponse), nil
	}
}
