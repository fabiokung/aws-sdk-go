// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package ssm provides a client for Amazon Simple Systems Management Service.
package ssm

import (
	"sync"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
)

var oprw sync.Mutex

// CreateAssociationRequest generates a request for the CreateAssociation operation.
func (c *SSM) CreateAssociationRequest(input *CreateAssociationInput) (req *aws.Request, output *CreateAssociationOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opCreateAssociation == nil {
		opCreateAssociation = &aws.Operation{
			Name:       "CreateAssociation",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &CreateAssociationInput{}
	}

	req = c.newRequest(opCreateAssociation, input, output)
	output = &CreateAssociationOutput{}
	req.Data = output
	return
}

// Associates the specified configuration document with the specified instance.
//
// When you associate a configuration document with an instance, the configuration
// agent on the instance processes the configuration document and configures
// the instance as specified.
//
// If you associate a configuration document with an instance that already
// has an associated configuration document, we replace the current configuration
// document with the new configuration document.
func (c *SSM) CreateAssociation(input *CreateAssociationInput) (output *CreateAssociationOutput, err error) {
	req, out := c.CreateAssociationRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreateAssociation *aws.Operation

// CreateAssociationBatchRequest generates a request for the CreateAssociationBatch operation.
func (c *SSM) CreateAssociationBatchRequest(input *CreateAssociationBatchInput) (req *aws.Request, output *CreateAssociationBatchOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opCreateAssociationBatch == nil {
		opCreateAssociationBatch = &aws.Operation{
			Name:       "CreateAssociationBatch",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &CreateAssociationBatchInput{}
	}

	req = c.newRequest(opCreateAssociationBatch, input, output)
	output = &CreateAssociationBatchOutput{}
	req.Data = output
	return
}

// Associates the specified configuration documents with the specified instances.
//
// When you associate a configuration document with an instance, the configuration
// agent on the instance processes the configuration document and configures
// the instance as specified.
//
// If you associate a configuration document with an instance that already
// has an associated configuration document, we replace the current configuration
// document with the new configuration document.
func (c *SSM) CreateAssociationBatch(input *CreateAssociationBatchInput) (output *CreateAssociationBatchOutput, err error) {
	req, out := c.CreateAssociationBatchRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreateAssociationBatch *aws.Operation

// CreateDocumentRequest generates a request for the CreateDocument operation.
func (c *SSM) CreateDocumentRequest(input *CreateDocumentInput) (req *aws.Request, output *CreateDocumentOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opCreateDocument == nil {
		opCreateDocument = &aws.Operation{
			Name:       "CreateDocument",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &CreateDocumentInput{}
	}

	req = c.newRequest(opCreateDocument, input, output)
	output = &CreateDocumentOutput{}
	req.Data = output
	return
}

// Creates a configuration document.
//
// After you create a configuration document, you can use CreateAssociation
// to associate it with one or more running instances.
func (c *SSM) CreateDocument(input *CreateDocumentInput) (output *CreateDocumentOutput, err error) {
	req, out := c.CreateDocumentRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreateDocument *aws.Operation

// DeleteAssociationRequest generates a request for the DeleteAssociation operation.
func (c *SSM) DeleteAssociationRequest(input *DeleteAssociationInput) (req *aws.Request, output *DeleteAssociationOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDeleteAssociation == nil {
		opDeleteAssociation = &aws.Operation{
			Name:       "DeleteAssociation",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &DeleteAssociationInput{}
	}

	req = c.newRequest(opDeleteAssociation, input, output)
	output = &DeleteAssociationOutput{}
	req.Data = output
	return
}

// Disassociates the specified configuration document from the specified instance.
//
// When you disassociate a configuration document from an instance, it does
// not change the configuration of the instance. To change the configuration
// state of an instance after you disassociate a configuration document, you
// must create a new configuration document with the desired configuration and
// associate it with the instance.
func (c *SSM) DeleteAssociation(input *DeleteAssociationInput) (output *DeleteAssociationOutput, err error) {
	req, out := c.DeleteAssociationRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteAssociation *aws.Operation

// DeleteDocumentRequest generates a request for the DeleteDocument operation.
func (c *SSM) DeleteDocumentRequest(input *DeleteDocumentInput) (req *aws.Request, output *DeleteDocumentOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDeleteDocument == nil {
		opDeleteDocument = &aws.Operation{
			Name:       "DeleteDocument",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &DeleteDocumentInput{}
	}

	req = c.newRequest(opDeleteDocument, input, output)
	output = &DeleteDocumentOutput{}
	req.Data = output
	return
}

// Deletes the specified configuration document.
//
// You must use DeleteAssociation to disassociate all instances that are associated
// with the configuration document before you can delete it.
func (c *SSM) DeleteDocument(input *DeleteDocumentInput) (output *DeleteDocumentOutput, err error) {
	req, out := c.DeleteDocumentRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteDocument *aws.Operation

// DescribeAssociationRequest generates a request for the DescribeAssociation operation.
func (c *SSM) DescribeAssociationRequest(input *DescribeAssociationInput) (req *aws.Request, output *DescribeAssociationOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDescribeAssociation == nil {
		opDescribeAssociation = &aws.Operation{
			Name:       "DescribeAssociation",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &DescribeAssociationInput{}
	}

	req = c.newRequest(opDescribeAssociation, input, output)
	output = &DescribeAssociationOutput{}
	req.Data = output
	return
}

// Describes the associations for the specified configuration document or instance.
func (c *SSM) DescribeAssociation(input *DescribeAssociationInput) (output *DescribeAssociationOutput, err error) {
	req, out := c.DescribeAssociationRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeAssociation *aws.Operation

// DescribeDocumentRequest generates a request for the DescribeDocument operation.
func (c *SSM) DescribeDocumentRequest(input *DescribeDocumentInput) (req *aws.Request, output *DescribeDocumentOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDescribeDocument == nil {
		opDescribeDocument = &aws.Operation{
			Name:       "DescribeDocument",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &DescribeDocumentInput{}
	}

	req = c.newRequest(opDescribeDocument, input, output)
	output = &DescribeDocumentOutput{}
	req.Data = output
	return
}

// Describes the specified configuration document.
func (c *SSM) DescribeDocument(input *DescribeDocumentInput) (output *DescribeDocumentOutput, err error) {
	req, out := c.DescribeDocumentRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeDocument *aws.Operation

// GetDocumentRequest generates a request for the GetDocument operation.
func (c *SSM) GetDocumentRequest(input *GetDocumentInput) (req *aws.Request, output *GetDocumentOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opGetDocument == nil {
		opGetDocument = &aws.Operation{
			Name:       "GetDocument",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &GetDocumentInput{}
	}

	req = c.newRequest(opGetDocument, input, output)
	output = &GetDocumentOutput{}
	req.Data = output
	return
}

// Gets the contents of the specified configuration document.
func (c *SSM) GetDocument(input *GetDocumentInput) (output *GetDocumentOutput, err error) {
	req, out := c.GetDocumentRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetDocument *aws.Operation

// ListAssociationsRequest generates a request for the ListAssociations operation.
func (c *SSM) ListAssociationsRequest(input *ListAssociationsInput) (req *aws.Request, output *ListAssociationsOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opListAssociations == nil {
		opListAssociations = &aws.Operation{
			Name:       "ListAssociations",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &ListAssociationsInput{}
	}

	req = c.newRequest(opListAssociations, input, output)
	output = &ListAssociationsOutput{}
	req.Data = output
	return
}

// Lists the associations for the specified configuration document or instance.
func (c *SSM) ListAssociations(input *ListAssociationsInput) (output *ListAssociationsOutput, err error) {
	req, out := c.ListAssociationsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListAssociations *aws.Operation

// ListDocumentsRequest generates a request for the ListDocuments operation.
func (c *SSM) ListDocumentsRequest(input *ListDocumentsInput) (req *aws.Request, output *ListDocumentsOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opListDocuments == nil {
		opListDocuments = &aws.Operation{
			Name:       "ListDocuments",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &ListDocumentsInput{}
	}

	req = c.newRequest(opListDocuments, input, output)
	output = &ListDocumentsOutput{}
	req.Data = output
	return
}

// Describes one or more of your configuration documents.
func (c *SSM) ListDocuments(input *ListDocumentsInput) (output *ListDocumentsOutput, err error) {
	req, out := c.ListDocumentsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListDocuments *aws.Operation

// UpdateAssociationStatusRequest generates a request for the UpdateAssociationStatus operation.
func (c *SSM) UpdateAssociationStatusRequest(input *UpdateAssociationStatusInput) (req *aws.Request, output *UpdateAssociationStatusOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opUpdateAssociationStatus == nil {
		opUpdateAssociationStatus = &aws.Operation{
			Name:       "UpdateAssociationStatus",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	if input == nil {
		input = &UpdateAssociationStatusInput{}
	}

	req = c.newRequest(opUpdateAssociationStatus, input, output)
	output = &UpdateAssociationStatusOutput{}
	req.Data = output
	return
}

// Updates the status of the configuration document associated with the specified
// instance.
func (c *SSM) UpdateAssociationStatus(input *UpdateAssociationStatusInput) (output *UpdateAssociationStatusOutput, err error) {
	req, out := c.UpdateAssociationStatusRequest(input)
	output = out
	err = req.Send()
	return
}

var opUpdateAssociationStatus *aws.Operation

// Describes an association of a configuration document and an instance.
type Association struct {
	// The ID of the instance.
	InstanceID *string `locationName:"InstanceId" type:"string"`

	// The name of the configuration document.
	Name *string `type:"string"`

	metadataAssociation `json:"-" xml:"-"`
}

type metadataAssociation struct {
	SDKShapeTraits bool `type:"structure"`
}

// Describes an association.
type AssociationDescription struct {
	// The date when the association was made.
	Date *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The ID of the instance.
	InstanceID *string `locationName:"InstanceId" type:"string"`

	// The name of the configuration document.
	Name *string `type:"string"`

	// The association status.
	Status *AssociationStatus `type:"structure"`

	metadataAssociationDescription `json:"-" xml:"-"`
}

type metadataAssociationDescription struct {
	SDKShapeTraits bool `type:"structure"`
}

// Describes a filter.
type AssociationFilter struct {
	// The name of the filter.
	Key *string `locationName:"key" type:"string" required:"true"`

	// The filter value.
	Value *string `locationName:"value" type:"string" required:"true"`

	metadataAssociationFilter `json:"-" xml:"-"`
}

type metadataAssociationFilter struct {
	SDKShapeTraits bool `type:"structure"`
}

// Describes an association status.
type AssociationStatus struct {
	// A user-defined string.
	AdditionalInfo *string `type:"string"`

	// The date when the status changed.
	Date *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`

	// The reason for the status.
	Message *string `type:"string" required:"true"`

	// The status.
	Name *string `type:"string" required:"true"`

	metadataAssociationStatus `json:"-" xml:"-"`
}

type metadataAssociationStatus struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreateAssociationBatchInput struct {
	// One or more associations.
	Entries []*CreateAssociationBatchRequestEntry `locationNameList:"entries" type:"list" required:"true"`

	metadataCreateAssociationBatchInput `json:"-" xml:"-"`
}

type metadataCreateAssociationBatchInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreateAssociationBatchOutput struct {
	// Information about the associations that failed.
	Failed []*FailedCreateAssociation `locationNameList:"FailedCreateAssociationEntry" type:"list"`

	// Information about the associations that succeeded.
	Successful []*AssociationDescription `locationNameList:"AssociationDescription" type:"list"`

	metadataCreateAssociationBatchOutput `json:"-" xml:"-"`
}

type metadataCreateAssociationBatchOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Describes the association of a configuration document and an instance.
type CreateAssociationBatchRequestEntry struct {
	// The ID of the instance.
	InstanceID *string `locationName:"InstanceId" type:"string"`

	// The name of the configuration document.
	Name *string `type:"string"`

	metadataCreateAssociationBatchRequestEntry `json:"-" xml:"-"`
}

type metadataCreateAssociationBatchRequestEntry struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreateAssociationInput struct {
	// The ID of the instance.
	InstanceID *string `locationName:"InstanceId" type:"string" required:"true"`

	// The name of the configuration document.
	Name *string `type:"string" required:"true"`

	metadataCreateAssociationInput `json:"-" xml:"-"`
}

type metadataCreateAssociationInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreateAssociationOutput struct {
	// Information about the association.
	AssociationDescription *AssociationDescription `type:"structure"`

	metadataCreateAssociationOutput `json:"-" xml:"-"`
}

type metadataCreateAssociationOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreateDocumentInput struct {
	// A valid JSON file. For more information about the contents of this file,
	// see Configuration Document (http://docs.aws.amazon.com/ssm/latest/APIReference/aws-ssm-document.html).
	Content *string `type:"string" required:"true"`

	// A name for the configuration document.
	Name *string `type:"string" required:"true"`

	metadataCreateDocumentInput `json:"-" xml:"-"`
}

type metadataCreateDocumentInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreateDocumentOutput struct {
	// Information about the configuration document.
	DocumentDescription *DocumentDescription `type:"structure"`

	metadataCreateDocumentOutput `json:"-" xml:"-"`
}

type metadataCreateDocumentOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteAssociationInput struct {
	// The ID of the instance.
	InstanceID *string `locationName:"InstanceId" type:"string" required:"true"`

	// The name of the configuration document.
	Name *string `type:"string" required:"true"`

	metadataDeleteAssociationInput `json:"-" xml:"-"`
}

type metadataDeleteAssociationInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteAssociationOutput struct {
	metadataDeleteAssociationOutput `json:"-" xml:"-"`
}

type metadataDeleteAssociationOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteDocumentInput struct {
	// The name of the configuration document.
	Name *string `type:"string" required:"true"`

	metadataDeleteDocumentInput `json:"-" xml:"-"`
}

type metadataDeleteDocumentInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteDocumentOutput struct {
	metadataDeleteDocumentOutput `json:"-" xml:"-"`
}

type metadataDeleteDocumentOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeAssociationInput struct {
	// The ID of the instance.
	InstanceID *string `locationName:"InstanceId" type:"string" required:"true"`

	// The name of the configuration document.
	Name *string `type:"string" required:"true"`

	metadataDescribeAssociationInput `json:"-" xml:"-"`
}

type metadataDescribeAssociationInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeAssociationOutput struct {
	// Information about the association.
	AssociationDescription *AssociationDescription `type:"structure"`

	metadataDescribeAssociationOutput `json:"-" xml:"-"`
}

type metadataDescribeAssociationOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeDocumentInput struct {
	// The name of the configuration document.
	Name *string `type:"string" required:"true"`

	metadataDescribeDocumentInput `json:"-" xml:"-"`
}

type metadataDescribeDocumentInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeDocumentOutput struct {
	// Information about the configuration document.
	Document *DocumentDescription `type:"structure"`

	metadataDescribeDocumentOutput `json:"-" xml:"-"`
}

type metadataDescribeDocumentOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// Describes a configuration document.
type DocumentDescription struct {
	// The date when the configuration document was created.
	CreatedDate *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The name of the configuration document.
	Name *string `type:"string"`

	// The SHA1 hash of the document, which you can use for verification purposes.
	SHA1 *string `locationName:"Sha1" type:"string"`

	// The status of the configuration document.
	Status *string `type:"string"`

	metadataDocumentDescription `json:"-" xml:"-"`
}

type metadataDocumentDescription struct {
	SDKShapeTraits bool `type:"structure"`
}

// Describes a filter.
type DocumentFilter struct {
	// The name of the filter.
	Key *string `locationName:"key" type:"string" required:"true"`

	// The value of the filter.
	Value *string `locationName:"value" type:"string" required:"true"`

	metadataDocumentFilter `json:"-" xml:"-"`
}

type metadataDocumentFilter struct {
	SDKShapeTraits bool `type:"structure"`
}

// Describes the name of a configuration document.
type DocumentIdentifier struct {
	// The name of the configuration document.
	Name *string `type:"string"`

	metadataDocumentIdentifier `json:"-" xml:"-"`
}

type metadataDocumentIdentifier struct {
	SDKShapeTraits bool `type:"structure"`
}

// Describes a failed association.
type FailedCreateAssociation struct {
	// The association.
	Entry *CreateAssociationBatchRequestEntry `type:"structure"`

	// The source of the failure.
	Fault *string `type:"string"`

	// A description of the failure.
	Message *string `type:"string"`

	metadataFailedCreateAssociation `json:"-" xml:"-"`
}

type metadataFailedCreateAssociation struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetDocumentInput struct {
	// The name of the configuration document.
	Name *string `type:"string" required:"true"`

	metadataGetDocumentInput `json:"-" xml:"-"`
}

type metadataGetDocumentInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetDocumentOutput struct {
	// The contents of the configuration document.
	Content *string `type:"string"`

	// The name of the configuration document.
	Name *string `type:"string"`

	metadataGetDocumentOutput `json:"-" xml:"-"`
}

type metadataGetDocumentOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListAssociationsInput struct {
	// One or more filters. Use a filter to return a more specific list of results.
	AssociationFilterList []*AssociationFilter `locationNameList:"AssociationFilter" type:"list" required:"true"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`

	metadataListAssociationsInput `json:"-" xml:"-"`
}

type metadataListAssociationsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListAssociationsOutput struct {
	// The associations.
	Associations []*Association `locationNameList:"Association" type:"list"`

	// The token to use when requesting the next set of items. If there are no additional
	// items to return, the string is empty.
	NextToken *string `type:"string"`

	metadataListAssociationsOutput `json:"-" xml:"-"`
}

type metadataListAssociationsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListDocumentsInput struct {
	// One or more filters. Use a filter to return a more specific list of results.
	DocumentFilterList []*DocumentFilter `locationNameList:"DocumentFilter" type:"list"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`

	metadataListDocumentsInput `json:"-" xml:"-"`
}

type metadataListDocumentsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListDocumentsOutput struct {
	// The names of the configuration documents.
	DocumentIdentifiers []*DocumentIdentifier `locationNameList:"DocumentIdentifier" type:"list"`

	// The token to use when requesting the next set of items. If there are no additional
	// items to return, the string is empty.
	NextToken *string `type:"string"`

	metadataListDocumentsOutput `json:"-" xml:"-"`
}

type metadataListDocumentsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type UpdateAssociationStatusInput struct {
	// The association status.
	AssociationStatus *AssociationStatus `type:"structure" required:"true"`

	// The ID of the instance.
	InstanceID *string `locationName:"InstanceId" type:"string" required:"true"`

	// The name of the configuration document.
	Name *string `type:"string" required:"true"`

	metadataUpdateAssociationStatusInput `json:"-" xml:"-"`
}

type metadataUpdateAssociationStatusInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type UpdateAssociationStatusOutput struct {
	// Information about the association.
	AssociationDescription *AssociationDescription `type:"structure"`

	metadataUpdateAssociationStatusOutput `json:"-" xml:"-"`
}

type metadataUpdateAssociationStatusOutput struct {
	SDKShapeTraits bool `type:"structure"`
}