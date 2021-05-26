# direktiv-actions-ghupload

This action uploads a namespace or workflow variable

## Usage

See [action.yaml](action.yml)

### Upload namespace variable as a file to direktiv

```yaml
steps:
  - name: get-something
    run: wget https://i.kym-cdn.com/entries/icons/original/000/022/017/thumb.png
  - name: upload
    id: upload 
    with:
      server: playground.direktiv.io
      namespace: trent
      variable: test3
      data: ./thumb.png
    uses: vorteil/direktiv-actions-ghupload@master
  - name: get output
    run: |
      echo ${{ steps.execute.outputs.body }}
```

### Upload json data as a workflow variable to direktiv

```yaml
steps:
  - name: upload
    id: upload 
    with:
      server: playground.direktiv.io
      namespace: trent
      variable: test2
      workflow: test
      data: |
        {
          "hello": "world"
        }
    uses: vorteil/direktiv-actions-ghupload@master
  - name: get output
    run: |
      echo ${{ steps.execute.outputs.body }}
```

### Using authentication token

```yaml
steps:
  - name: get-something
    run: wget https://i.kym-cdn.com/entries/icons/original/000/022/017/thumb.png
  - name: upload
    id: upload 
    with:
      server: playground.direktiv.io
      namespace: trent
      variable: test3
      data: ./thumb.png
      token: ${{ secrets.DIREKTIV_TOKEN }}
    uses: vorteil/direktiv-actions-ghupload@master
  - name: get output
    run: |
      echo ${{ steps.execute.outputs.body }}
```
