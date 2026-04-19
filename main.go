package main

import (
	"context"
	"fmt"
	"os"
	"sync"

)

type VolumeResult struct {

	AccountID string
	VolumeID string
	Size string
	State string
	Region string
}

func main(){

	//Define the AWS Profiles/Accounts you want to scan
	profiles := []string{"default", "production", "staging"}
	region := "us-east-1"

	var wg sync.WaitGroup

	resultsChan := make(chan []VolumeResult, len(profiles))

	pterm.DefaultHeader.WithFullWidth().Println("EBS Waste finder")

	spinner, _ := pterm.DefaultSpinner.Start("SCanning AWS accounts")

	for _, profile := range profiles{
		wg.Add(1)
		go func(p string){
			defer wg.Done()
			resultsChan <- findUnusedVolumes(p, region)
		}(profile)
	}

	go func(){
		wg.Wait()
		close(resultsChan)
	}()

	var allVolumes []VolumeResult
	for res := range resultsChan {
		allVolumes = append(allVolumes, res...)
	}

	spinner.Success("scan complete!")
	renderTable(allVolumes)
}

func findUnusedVolumes(profile string, region string) []VolumeResult{
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx,
					config.WithSharedConfigProfile(profile),
				config.WithRegion(region),
			)

	if err != nil {
		return nil
	}

	client := ec2.NewFromConfig(cfg)

	input := &ec2.DescribeVolumesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("status"),
				Values: []string{"available"},
			},
		},
	}

	resp, err := client.DescribeVolumes(ctx, input)
	if err != nil {
		return nil
	}

	var results []VolumeResult
	for _, v := range resp.Volumes {
		results = append(results, VolumeResult{
			AccountID: profile, // Using profile name as identifier
			VolumeID:  *v.VolumeId,
			Size:      *v.Size,
			State:     string(v.State),
			Region:    region,
		})
	}
	return results
}

func renderTable(volumes []VolumeResult) {
	tableData := pterm.TableData{
		{"Account/Profile", "Volume ID", "Size (GB)", "State", "Region"},
	}

	for _, v := range volumes {
		tableData = append(tableData, []string{
			v.AccountID,
			v.VolumeID,
			fmt.Sprintf("%d", v.Size),
			pterm.LightRed(v.State),
			v.Region,
		})
	}

	pterm.DefaultTable.WithHasSeparator().WithData(tableData).Render()
}