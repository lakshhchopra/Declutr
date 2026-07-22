import json
import requests
from typing import Dict, Any, Optional

class DeclutrClient:
    """Official Python SDK Client for Declutr Developer Platform."""

    def __init__(self, api_key: Optional[str] = None, access_token: Optional[str] = None, base_url: str = "http://localhost:8080"):
        self.api_key = api_key
        self.access_token = access_token
        self.base_url = base_url.rstrip("/")

    def _get_headers(self) -> Dict[str, str]:
        headers = {"Content-Type": "application/json"}
        if self.api_key:
            headers["Authorization"] = f"Bearer {self.api_key}"
        elif self.access_token:
            headers["Authorization"] = f"Bearer {self.access_token}"
        return headers

    def _request(self, method: str, endpoint: str, json_data: Optional[Dict[str, Any]] = None) -> Dict[str, Any]:
        url = f"{self.base_url}{endpoint}"
        resp = requests.request(method, url, headers=self._get_headers(), json=json_data)
        if resp.status_code >= 400:
            raise RuntimeError(f"Declutr API Error [{resp.status_code}]: {resp.text}")
        return resp.json()

    def search(self, query: str, filters: Optional[Dict[str, Any]] = None) -> Dict[str, Any]:
        """Execute hybrid search across user/organization vaults."""
        return self._request("POST", "/api/v1/search/query", {"query": query, "filters": filters or {}})

    def chat(self, conversation_id: str, message: str) -> Dict[str, Any]:
        """Send prompt to RAG Grounded AI Copilot."""
        return self._request("POST", "/api/v1/copilot/messages", {"conversation_id": conversation_id, "content": message})

    def run_workflow(self, workflow_id: str, input_data: Optional[Dict[str, Any]] = None) -> Dict[str, Any]:
        """Trigger workflow automation execution."""
        return self._request("POST", "/api/v1/workflows/run", {"workflow_id": workflow_id, "input": input_data or {}})
