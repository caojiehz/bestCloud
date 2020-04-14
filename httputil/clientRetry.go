package httputil

import (
	"encoding/json"
	"errors"
)

func RetryGet(paras GetParasTuple) (err error) {
	var data []byte
	for i := 0; i < paras.RetryNum; i++ {
		if data, err = Get(paras); err == nil {
			if paras.Resp == nil {
				return nil
			}
			if err = json.Unmarshal(data, paras.Resp); err == nil {
				return nil
			}
		}
	}
	return err
}

func RetryGetForm(paras GetFormParasTuple) (err error) {
	var data []byte
	for i := 0; i < paras.RetryNum; i++ {
		if data, err = GetForm(paras); err == nil {
			if paras.Resp == nil {
				return nil
			}
			if err = json.Unmarshal(data, paras.Resp); err == nil {
				return nil
			}
		}
	}
	return err
}

func RetryPost(paras PostParasTuple) (err error) {
	var data []byte
	for i := 0; i < paras.RetryNum; i++ {
		if data, err = Post(paras); err == nil {
			if paras.Resp == nil {
				return nil
			}
			if err = json.Unmarshal(data, paras.Resp); err == nil {
				return nil
			}
		}
	}
	return err
}

func RetryPostForm(paras PostFormParasTuple) (err error) {
	var data []byte
	for i := 0; i < paras.RetryNum; i++ {
		if data, err = PostForm(paras); err == nil {
			if paras.Resp == nil {
				return nil
			}
			if err = json.Unmarshal(data, paras.Resp); err == nil {
				return nil
			}
		}
	}
	return err
}

func RetryPostBytes(paras PostBytesTuple) (data []byte, err error) {
	for i := 0; i < paras.RetryNum; i++ {
		if data, err := PostBytes(paras); err == nil {
			return data, nil
		}
	}
	return nil, errors.New("retry post bytes fail")
}

func RetryDelete(paras GetParasTuple) (err error) {
	var data []byte
	for i := 0; i < paras.RetryNum; i++ {
		if data, err = Delete(paras); err == nil {
			if paras.Resp == nil {
				return nil
			}
			if err = json.Unmarshal(data, paras.Resp); err == nil {
				return nil
			}
		}
	}
	return err
}
