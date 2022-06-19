package service

import (
	"agent-app/dto"
	"agent-app/repository"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type OfferService struct {
	CompanyRepo repository.ICompanyRepository
}

type IOfferService interface {
	Add(*dto.JobOfferRequestDTO, string) (*dto.JobOfferResponseDTO, error)
	GetCompanysOffers(int) ([]*dto.JobOfferResponseDTO, error)
	GetAll() ([]*dto.JobOfferResponseDTO, error)
	Search(string) ([]*dto.JobOfferResponseDTO, error)
	GetJobOfferById(int) (*dto.JobOfferResponseDTO, error)
	DeleteJobOffer(int, string) error
}

func NewOfferService(companyRepository repository.ICompanyRepository) IOfferService {
	return &OfferService{
		companyRepository,
	}
}

func (service *OfferService) Add(offer *dto.JobOfferRequestDTO, ownerAuth0ID string) (*dto.JobOfferResponseDTO, error) {
	err := offer.Validate()
	if err != nil {
		return nil, err
	}

	company, err := service.CompanyRepo.GetByID(offer.CompanyID)

	if err != nil {
		return nil, err
	}

	if ownerAuth0ID != company.OwnerAuth0ID {
		return nil, errors.New("Only company owner is allowed to publish job offers")
	}

	b, _ := json.Marshal(&offer)
	endpoint := os.Getenv("JOBS_MS")
	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(b))
	req.Header.Set("content-type", "application/json")
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 201 {
		fmt.Println(res.StatusCode)
		b, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(b))
		return nil, errors.New("Failed to create job offer")
	}

	body, _ := io.ReadAll(res.Body)
	var forReturn dto.JobOfferResponseDTO
	json.Unmarshal(body, &forReturn)
	return &forReturn, nil
}

func (service *OfferService) GetAll() ([]*dto.JobOfferResponseDTO, error) {
	endpoint := os.Getenv("JOBS_MS")
	req, _ := http.NewRequest("GET", endpoint, nil)
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		fmt.Println(res.StatusCode)
		b, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(b))
		return nil, errors.New("Failed to get job offers")
	}

	body, _ := io.ReadAll(res.Body)
	var forReturn []*dto.JobOfferResponseDTO
	json.Unmarshal(body, &forReturn)
	return forReturn, nil
}

func (service *OfferService) Search(param string) ([]*dto.JobOfferResponseDTO, error) {
	endpoint := os.Getenv("JOBS_MS")
	withParam := endpoint + "/search?param=" + param
	req, _ := http.NewRequest("GET", withParam, nil)
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		fmt.Println(res.StatusCode)
		b, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(b))
		return nil, errors.New("Failed to search job offers")
	}

	body, _ := io.ReadAll(res.Body)
	var forReturn []*dto.JobOfferResponseDTO
	json.Unmarshal(body, &forReturn)
	return forReturn, nil
}

func (service *OfferService) GetCompanysOffers(id int) ([]*dto.JobOfferResponseDTO, error) {
	endpoint := os.Getenv("JOBS_MS")
	withParam := endpoint + "/company/" + strconv.Itoa(id)
	req, _ := http.NewRequest("GET", withParam, nil)
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		fmt.Println(res.StatusCode)
		b, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(b))
		return nil, errors.New("Failed to get job offers for company")
	}

	body, _ := io.ReadAll(res.Body)
	var forReturn []*dto.JobOfferResponseDTO
	json.Unmarshal(body, &forReturn)
	return forReturn, nil
}

func (service *OfferService) GetJobOfferById(id int) (*dto.JobOfferResponseDTO, error) {
	endpoint := os.Getenv("JOBS_MS")
	withParam := endpoint + "/" + strconv.Itoa(id)
	req, _ := http.NewRequest("GET", withParam, nil)
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		fmt.Println(res.StatusCode)
		b, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(b))
		return nil, errors.New(fmt.Sprintf("Failed to get job offer with id: %d", id))
	}

	body, _ := io.ReadAll(res.Body)
	var forReturn *dto.JobOfferResponseDTO
	json.Unmarshal(body, &forReturn)
	return forReturn, nil
}

func (service *OfferService) DeleteJobOffer(id int, ownerID string) error {
	endpoint := os.Getenv("JOBS_MS")
	withParam := endpoint + "/" + strconv.Itoa(id)
	req, _ := http.NewRequest("DELETE", withParam, nil)
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))

	existingOffer, err := service.GetJobOfferById(id)

	if err != nil {
		return err
	}

	company, err := service.CompanyRepo.GetByID(existingOffer.CompanyID)

	if err != nil {
		return err
	}

	if ownerID != company.OwnerAuth0ID {
		return errors.New("Only company owner is allowed to delete job offers")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 204 {
		fmt.Println(res.StatusCode)
		b, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(b))
		return errors.New(string(b))
	}

	return nil
}
